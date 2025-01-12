package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"TTMS/kitex_gen/order"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
+---------------------------------------------------+
|					delayQueue->member				|
+---------------------------------------------------+
|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|
+---------------------------------------------------+
*/
var redisC *redis.Client
var delayQueue = "delay_queue"
var targetQueue = "target_queue"

func init() {
	redisC = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddress,
		Password: consts.RedisPassword,
		DB:       consts.RedisTicketDB,
	})
	ctx := context.Background()
	go toTargetQueue(ctx)
	go eventLoop(ctx)
}

// ToDelayQueue 将任务添加到延迟队列
func ToDelayQueue(ctx context.Context, orderInfo string, timeUnix float64) {
	err := redisC.ZAdd(ctx, delayQueue, redis.Z{Member: orderInfo, Score: timeUnix}).Err()
	log.Println("ToDelayQueue time = ", time.Unix(int64(timeUnix), 0).Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("ToDelayQueue ", err)
	}
}

// RemoveFromDelayQueue 按时支付订单，从延时队列中取出该订单
func RemoveFromDelayQueue(ctx context.Context, req *order.CommitOrderRequest) error {
	//TODO 修改逻辑，保证幂等 1. 先查看redis中有没有，如果没有，说明订单已经过期，直接返回；2. 调用支付接口；3. 若用户成功支付，将order信息从zSet中删除并修改数据库、若用户支付失败，则不进行任何操作
	// 第三步redis和MySQL不能保证原子性，这里我们只要保证MySQL状态更新成功就可以了。
	// 因为就算redis删除失败，在之后订单自动过期的逻辑里，我们也会查看MySQL中订单状态，如果已经支付成功。是不会让订单被退回的。
	orderInfo := fmt.Sprintf("%d;%d;%d;%d", req.UserId, req.ScheduleId, req.SeatRow, req.SeatCol)
	logrus.Debug("orderInfo = ", orderInfo)
	score, err := redisC.ZScore(ctx, delayQueue, orderInfo).Result()
	logrus.Debug("score = ", score, " err = ", err)
	if err != nil {
		logrus.Error("failed to get order from delay queue:", err)
		return err
	}
	// 模拟调用第三方支付接口，指定支付过期时间
	time.Sleep(time.Second)
	if time.Now().Unix() > int64(score) { //或者第三方接口返回支付过期，支付失败
		return errors.New("订单已经过期")
	}
	// 按时支付成功，从延迟队列中删除订单
	count, err := redisC.ZRem(ctx, delayQueue, orderInfo).Result()
	logrus.Debug("count = ", count, " err = ", err)
	if err != nil {
		logrus.Error("failed to remove order from delay queue:", err)
		return err
	}
	if count == 0 {
		return errors.New("订单已经过期")
	}
	//更改订单状态为已支付
	err = dao.UpdateOrder(req.UserId, req.ScheduleId, req.SeatRow, req.SeatCol, 1, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// toTargetQueue 处理延迟队列取出来的订单
func toTargetQueue(ctx context.Context) {
	for {
		// 从延迟队列中取出最小的时间戳，延迟3s,避免极端情况下，用户卡点支付，导致订单被错误地取消
		result, err := redisC.ZRangeByScoreWithScores(ctx, delayQueue, &redis.ZRangeBy{
			Min:    "-inf",
			Max:    fmt.Sprintf("%d", time.Now().Add(time.Second*3).Unix()),
			Offset: 0,
			Count:  100,
		}).Result()
		if err != nil {
			// 处理错误
			panic(err)
		}

		if len(result) == 0 {
			// 延迟队列中没有数据，等待一段时间后再次查询
			time.Sleep(time.Second)
			continue
		}

		for _, z := range result {
			// 从延迟队列中移除该数据
			_, err := redisC.ZRem(ctx, delayQueue, z.Member).Result()
			if err != nil {
				// 处理错误
				panic(err)
			}

			// 将数据从延迟队列转移到目标队列
			logrus.Debug(z.Member, z.Score)
			_, err = redisC.LPush(ctx, targetQueue, z.Member).Result()
			if err != nil {
				// 处理错误
				panic(err)
			}
		}
	}
}

// 循环处理targetQueue中的数据 (这些都是超时未支付的订单，需要从数据库中删除order实体，重新放出票)
func eventLoop(ctx context.Context) {
	for {
		results, err := redisC.BLPop(ctx, time.Second*5, targetQueue).Result()
		if errors.Is(err, redis.Nil) {
			continue
		}
		logrus.Debug("now = ", time.Now().Format(time.DateTime))
		logrus.Debug("results = ", results)
		if len(results) == 0 {
			continue
		}
		results = results[1:]
		logrus.Debug("results = ", results)
		for _, result := range results {
			fmt.Println("result = ", result)
			data := strings.Split(result, ";")
			fmt.Println("data = ", data)
			d0, _ := strconv.Atoi(data[0])
			d1, _ := strconv.Atoi(data[1])
			d2, _ := strconv.Atoi(data[2])
			d3, _ := strconv.Atoi(data[3])
			//向ticket服务提交，更改票状态为 待售
			sendMessage(kafka.Message{Value: []byte(fmt.Sprintf("%d;%d;%d", d1, d2, d3))})
			err = dao.DeleteOrder(ctx, d0, d1, d2, d3)
			if err != nil {
				logrus.Error("failed to delete order:", err)
			}
		}

	}
}
