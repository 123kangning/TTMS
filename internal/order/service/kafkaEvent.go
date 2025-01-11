package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func NewKafkaEventLoop() {
	go orderBuy()
	logrus.Info("kafka event start")
}

func orderBuy() {
	tag := "orderBuy"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{consts.KafkaBroker1Url, consts.KafkaBroker2Url, consts.KafkaBroker3Url},
		GroupID: consts.OrderGroupID,
		Topic:   consts.OrderBuyTopic,
	})
	defer r.Close()

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			logrus.Error(tag, "failed to fetch message:", err)
			break
		}
		err = AddOrderHandler(m)
		//未成功处理消息，暂时不做处理
		if err != nil {
			logrus.Error(tag, "failed to handle message:", err)
			continue
		}
		//成功处理消息，提交消息
		if err := r.CommitMessages(ctx, m); err != nil {
			logrus.Error(tag, "failed to commit message:", err)
		}
	}
}

func AddOrderHandler(msg kafka.Message) error {
	data := strings.Split(string(msg.Value), ";")
	logrus.Debug("data = ", data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	d3, _ := strconv.Atoi(data[3])
	d5, _ := strconv.Atoi(data[5])

	ctx := context.Background()
	oldTicketCount := dao.GetOrderCount(ctx, d1, d2, d3, []int{1, 2})
	if oldTicketCount > 0 {
		//终止此消息，再次保证了一张票只能被一个人买到
		logrus.Error("重复消息，不进行处理")
		return nil
	}

	t := float64(time.Now().Add(consts.OrderDelayTime).Unix())
	//fmt.Println("time = ", time.Now().Format("2006-01-02 15:04:05"), " unix = ", time.Now().Unix())
	//fmt.Println("time = ", time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"), " unix = ", t)
	orderInfo := strings.Join(data[:4], ";")
	//fmt.Println("orderInfo = ", orderInfo)
	ToDelayQueue(ctx, orderInfo, t)
	//fmt.Println(d0, d1, d2, d3, data[4])
	err := dao.AddOrder(d0, d1, d2, d3, data[4], d5)
	//暂时未出现这种情况，先不管
	if err != nil {
		logrus.Error("dao.AddOrder failed,err=", err, "msg is", string(msg.Value))
		return err
	}

	//业务消费消息完成之后，再返回commit告知消息队列跳过该消息，获取下一条消息
	return nil
}
