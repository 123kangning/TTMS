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
	"TTMS/internal/order/dao"
	"TTMS/kitex_gen/order"
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/play/playservice"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var playClient playservice.Client

func InitPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := playservice.NewClient(
		consts.PlayServiceName,
		//client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 1, MaxIdlePerAddress: 100, MaxIdleGlobal: 10000, MaxIdleTimeout: time.Minute}),
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
	playClient = c
}

func GetAllOrderService(ctx context.Context, req *order.GetAllOrderRequest) (resp *order.GetAllOrderResponse, err error) {
	resp = &order.GetAllOrderResponse{BaseResp: &order.BaseResp{}}
	resp.List, _, err = dao.GetAllOrder(ctx, int(req.UserId), int(req.OrderType))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func GetOrderAnalysisService(ctx context.Context, req *order.GetOrderAnalysisRequest) (resp *order.GetOrderAnalysisResponse, err error) {
	//通过RPC 根据playID查找scheduleId,再根据scheduleId统计票房
	fmt.Println("通过RPC 根据playID查找scheduleId,再根据scheduleId统计票房")
	resp = &order.GetOrderAnalysisResponse{BaseResp: &order.BaseResp{}}
	resp1, _ := playClient.PlayToSchedule(ctx, &play.PlayToScheduleRequest{Id: req.PlayId})
	if resp1.BaseResp.StatusCode == 1 {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = resp1.BaseResp.StatusMessage
		log.Println(resp1.BaseResp.StatusMessage)
		return resp, nil
	}
	o := &order.OrderAnalysis{
		PlayId:       resp1.Play.Id,
		PlayName:     resp1.Play.Name,
		Price:        int32(resp1.Play.Price),
		PlayArea:     resp1.Play.Area,
		PlayDuration: resp1.Play.Duration,
		StartData:    resp1.Play.StartDate,
		EndData:      resp1.Play.EndDate,
	}
	o.TotalTicket, o.Sales, err = dao.GetOrderAnalysis(ctx, resp1.ScheduleList)
	resp.OrderAnalysis = o
	fmt.Println("o.TotalTicket = ", o.TotalTicket, " err = ", err)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	fmt.Println("resp = ", resp)
	return resp, nil
}
func CommitOrderService(ctx context.Context, req *order.CommitOrderRequest) (resp *order.CommitOrderResponse, err error) {
	resp = &order.CommitOrderResponse{BaseResp: &order.BaseResp{}}
	err = RemoveFromDelayQueue(ctx, req)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}

	return resp, nil
}
func UpdateOrderService(ctx context.Context, req *order.UpdateOrderRequest) (resp *order.UpdateOrderResponse, err error) {
	resp = &order.UpdateOrderResponse{BaseResp: &order.BaseResp{}}
	err = dao.UpdateOrder(req.UserId, req.ScheduleId, req.SeatRow, req.SeatCol, -1, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}

	return resp, nil
}
