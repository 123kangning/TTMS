package service

/*
+---------------------------------------------------------------------------+---------------------------------------------------------------+
|							BuyTicketMsg									|						ReturnTicketMsg							|
|---------------------------------------------------------------------------+---------------------------------------------------------------+
|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|	Time	|	Price	|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|	Time	|
+---------------------------------------------------------------------------+---------------------------------------------------------------+
*/
import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"TTMS/internal/ticket/redis"
	"TTMS/kitex_gen/order"
	"TTMS/kitex_gen/order/orderservice"
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/play/playservice"
	"TTMS/kitex_gen/ticket"
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var playClient playservice.Client
var orderClient orderservice.Client
var Loc *time.Location

func init() {
	var err error
	Loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("Error loading location:", err)
		return
	}
}

func OrderPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := orderservice.NewClient(
		consts.OrderServiceName,
		//client.WithMiddleware(mw.CommonMiddleware),
		//client.WithInstanceMW(mw.ClientMiddleware),
		// client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(consts.RPCTimeout),          // rpc timeout
		client.WithConnectTimeout(consts.ConnectTimeout),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	orderClient = c
}
func InitPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := playservice.NewClient(
		consts.PlayServiceName,
		//client.WithMiddleware(mw.CommonMiddleware),
		//client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(consts.RPCTimeout),          // rpc timeout
		client.WithConnectTimeout(consts.ConnectTimeout),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	playClient = c
}

func BatchAddTicketService(ctx context.Context, req *ticket.BatchAddTicketRequest) (resp *ticket.BatchAddTicketResponse, err error) {
	//fmt.Println(req.ScheduleId, req.Price, req.PlayName, req.StudioId, req.List)
	err = dao.BatchAddTicket(ctx, req.ScheduleId, req.Price, req.PlayName, req.StudioId, req.List)
	for _, s := range req.List {
		redis.AddTicket(int(req.ScheduleId), int(s.Row), int(s.Col), req.Price)
	}
	resp = &ticket.BatchAddTicketResponse{BaseResp: &ticket.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func UpdateTicketService(ctx context.Context, req *ticket.UpdateTicketRequest) (resp *ticket.UpdateTicketResponse, err error) {
	err = dao.UpdateTicket(ctx, req.ScheduleId, req.SeatRow, req.SeatCol, req.Price, req.Status)
	resp = &ticket.UpdateTicketResponse{BaseResp: &ticket.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	fmt.Println("resp = ", resp)
	return resp, nil
}

func GetAllTicketService(ctx context.Context, req *ticket.GetAllTicketRequest) (resp *ticket.GetAllTicketResponse, err error) {
	resp = &ticket.GetAllTicketResponse{BaseResp: &ticket.BaseResp{}}
	resp.List, err = dao.GetAllTicket(ctx, req.ScheduleId)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func TicketIsExist(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) (bool, error, string) {
	key := fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol)
	result, err := redis.TicketIsExist(key) //只查看票的状态，不抢票

	//redis没找到，去mysql再找一下，防止redis挂了，导致所有用户买不了票
	//TODO 如果后续新架构要支持防止缓存穿透的场景，应该将‘票未存在于redis中’这种情况直接拒绝。目前用的老架构，票会定时过期，所以还不能改。
	if err != nil {
		t := dao.GetTicket(ctx, ScheduleId, SeatRow, SeatCol)
		if t.Id > 0 && t.Status == 0 {
			return true, err, "mysql"
		}
	}
	if result {
		return true, err, "redis"
	}
	return false, err, ""
}
func BuyTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32, source string) {
	if source == "redis" {
		redis.BuyTicket(ctx, fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol))
	}
	//source=="mysql"，无论是否更新redis，mysql是一定要更新的
	/*考虑一个问题（前提：redis的票过期后）：
	redis分布式锁被释放前，异步更新db能否完成，若不能完成就会有超卖的风险
	直接用同步操作也可以，就是接口速度会慢一些，但是放心*/
	dao.BuyTicket(ctx, ScheduleId, SeatRow, SeatCol)

}
func BuyTicketService(ctx context.Context, req *ticket.BuyTicketRequest) (resp *ticket.BuyTicketResponse, err error) {
	resp = &ticket.BuyTicketResponse{BaseResp: &ticket.BaseResp{}}
	//查看票是否还存在（没有被别人买）
	key := fmt.Sprintf("%d;%d;%d", req.ScheduleId, req.SeatRow, req.SeatCol)

	result, err, source := TicketIsExist(ctx, req.ScheduleId, req.SeatRow, req.SeatCol) //只查看票的状态，不抢票
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, nil
	}

	//票已经被买,或者是同时抢票但没有抢到分布式锁
	if !result || !redis.AcquireLock(fmt.Sprintf("lock;%s", key)) {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("票已经被买").Error()
		return resp, nil
	}
	defer redis.ReleaseLock(fmt.Sprintf("lock;%s", key)) //释放锁

	//判断是否为有效的’买票时间‘,多人抢票时，只让抢到分布式锁的用户进行时间检查
	schedule, err := playClient.GetSchedule(ctx, &play.GetScheduleRequest{Id: req.ScheduleId})

	deadline, _ := time.Parse("2006-01-02 15:04:05", schedule.Schedule.ShowTime)
	deadline = deadline.In(Loc).Add(-8 * time.Hour)

	//log.Println("now = ", time.Now().Format("2006-01-02 15:04:05"))
	//log.Println("showtime = ", deadline)
	//log.Println("until = ", time.Until(deadline))
	if time.Until(deadline) < 10*time.Minute { //距离开场已经不足10分钟，停止售票
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("已停止售票").Error()
		return resp, nil
	}

	//抢到分布式锁,执行买票流程
	//选择是否更新redis，并强制更新mysql
	BuyTicket(ctx, req.ScheduleId, req.SeatRow, req.SeatCol, source)

	//成功抢到票,发送创建订单消息
	t := time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println("time = ", t)
	err = sendMessage(kafka.Message{
		Value: []byte(fmt.Sprintf("%d;%s;%s;%s", req.UserId, key, t,
			redis.GetTicketPrice(ctx, fmt.Sprintf("%d;price", req.ScheduleId))))})

	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func sendMessage(msg kafka.Message) error {
	w := &kafka.Writer{
		Addr:         kafka.TCP(consts.KafkaBroker1Url, consts.KafkaBroker2Url, consts.KafkaBroker3Url),
		Topic:        consts.OrderBuyTopic,
		Balancer:     &kafka.LeastBytes{}, //Hash, RoundRobin, LeastBytes
		RequiredAcks: kafka.RequireOne,    //Leader写入成功即可
	}
	defer func() {
		if err := w.Close(); err != nil {
			logrus.Error("failed to close writer:", err)
		}
	}()

	err := w.WriteMessages(context.Background(), msg)
	if err != nil {
		logrus.Error(w.Topic, "failed to write messages:", err)
		return err
	}
	return nil
}

func ReturnTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) {
	redis.ReturnTicket(ctx, fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol))
	go dao.ReturnTicket(ctx, ScheduleId, SeatRow, SeatCol)
}
func ReturnTicketService(ctx context.Context, req *ticket.ReturnTicketRequest) (resp *ticket.ReturnTicketResponse, err error) {
	//先判断是否为有效的’退票时间‘
	schedule, err := playClient.GetSchedule(ctx, &play.GetScheduleRequest{Id: req.ScheduleId})
	resp = &ticket.ReturnTicketResponse{BaseResp: &ticket.BaseResp{}}

	deadline, _ := time.Parse("2006-01-02 15:04:05", schedule.Schedule.ShowTime)
	deadline = deadline.In(Loc).Add(-8 * time.Hour)
	if time.Until(deadline) < 0 { //演出已经开始，停止退票
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("已停止退票").Error()
		return resp, nil
	}

	resp1, _ := orderClient.UpdateOrder(ctx, &order.UpdateOrderRequest{UserId: req.UserId, ScheduleId: req.ScheduleId,
		SeatRow: req.SeatRow, SeatCol: req.SeatCol})
	if resp1.BaseResp.StatusCode == 1 {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = resp1.BaseResp.StatusMessage
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	ReturnTicket(ctx, req.ScheduleId, req.SeatRow, req.SeatCol)
	return resp, nil
}
