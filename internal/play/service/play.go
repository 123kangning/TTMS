package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/play/dao"
	"TTMS/internal/play/redis"
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/studio"
	"TTMS/kitex_gen/studio/studioservice"
	"TTMS/kitex_gen/ticket"
	"TTMS/kitex_gen/ticket/ticketservice"
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"time"
)

var studioClient studioservice.Client
var ticketClient ticketservice.Client

func InitStudioRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := studioservice.NewClient(
		consts.StudioServiceName,
		//client.WithMiddleware(mw.CommonMiddleware),
		//client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	studioClient = c
}
func InitTicketRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := ticketservice.NewClient(
		consts.TicketServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	ticketClient = c
}
func AddPlayService(ctx context.Context, req *play.AddPlayRequest) (resp *play.AddPlayResponse, err error) {
	PlayInfo := &play.Play{Name: req.Name, Type: req.Type, Area: req.Area,
		Rating: req.Rating, Duration: req.Duration, StartDate: req.StartDate, EndDate: req.EndDate, Price: req.Price}
	err = dao.AddPlay(ctx, PlayInfo)
	resp = &play.AddPlayResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func UpdatePlayService(ctx context.Context, req *play.UpdatePlayRequest) (resp *play.UpdatePlayResponse, err error) {
	PlayInfo := &play.Play{Id: req.Id, Name: req.Name, Type: req.Type, Area: req.Area,
		Rating: req.Rating, Duration: req.Duration, StartDate: req.StartDate, EndDate: req.EndDate, Price: req.Price}
	err = dao.UpdatePlay(ctx, PlayInfo)
	resp = &play.UpdatePlayResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
		if req.Price != 0 {
			_, schList, _ := dao.PlayToSchedule(ctx, req.Id)
			redis.UpdatePlayPrice(ctx, schList, int(req.Price))
		}
	}
	return resp, nil
}

func DeletePlayService(ctx context.Context, req *play.DeletePlayRequest) (resp *play.DeletePlayResponse, err error) {
	err = dao.DeletePlay(ctx, req.Id)
	resp = &play.DeletePlayResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func GetAllPlayService(ctx context.Context, req *play.GetAllPlayRequest) (resp *play.GetAllPlayResponse, err error) {
	resp = &play.GetAllPlayResponse{BaseResp: &play.BaseResp{}}
	resp.List, err = dao.GetAllPlay(ctx, int(req.Current), int(req.PageSize))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func AddScheduleService(ctx context.Context, req *play.AddScheduleRequest) (resp *play.AddScheduleResponse, err error) {
	resp = &play.AddScheduleResponse{BaseResp: &play.BaseResp{}}
	resp0, err := studioClient.GetStudio(ctx, &studio.GetStudioRequest{Id: req.StudioId})
	if resp0.Result.Id == 0 { //判断演出厅是否存在
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("计划中的演出厅不存在").Error()
		return resp, nil
	}
	SInfo := &play.Schedule{PlayId: req.PlayId, StudioId: req.StudioId, ShowTime: req.ShowTime}
	id, err := dao.AddSchedule(ctx, SInfo)

	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, nil
	}
	re, _ := studioClient.GetAllSeat(ctx, &studio.GetAllSeatRequest{StudioId: req.StudioId, Current: 0, PageSize: 1000})
	p, err := dao.GetPlayById(req.PlayId)
	//log.Println(re.List)
	_, err = ticketClient.BatchAddTicket(ctx, &ticket.BatchAddTicketRequest{ScheduleId: id, StudioId: req.StudioId, Price: int32(p.Price), PlayName: p.Name, List: re.List})
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"

	}
	return resp, nil
}

func UpdateScheduleService(ctx context.Context, req *play.UpdateScheduleRequest) (resp *play.UpdateScheduleResponse, err error) {
	SInfo := &play.Schedule{Id: req.Id, PlayId: req.PlayId, StudioId: req.StudioId, ShowTime: req.ShowTime}
	err = dao.UpdateSchedule(ctx, SInfo)
	resp = &play.UpdateScheduleResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func DeleteScheduleService(ctx context.Context, req *play.DeleteScheduleRequest) (resp *play.DeleteScheduleResponse, err error) {
	err = dao.DeleteSchedule(ctx, req.Id)
	resp = &play.DeleteScheduleResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func GetAllScheduleService(ctx context.Context, req *play.GetAllScheduleRequest) (resp *play.GetAllScheduleResponse, err error) {
	resp = &play.GetAllScheduleResponse{BaseResp: &play.BaseResp{}}
	resp.List = make([]*play.Result, 0, req.PageSize)
	schedules, err := dao.GetAllSchedule(ctx, int(req.Current), int(req.PageSize))
	log.Println("schedule = ", schedules)
	for i, sch := range schedules {
		p, _ := dao.GetPlayById(sch.PlayId)
		resp1, err1 := studioClient.GetStudio(ctx, &studio.GetStudioRequest{Id: sch.StudioId})
		log.Println("play = ", p)
		log.Println("studio = ", resp1.Result)
		log.Println("err1 = ", err1)
		//log.Println("i = ", i, "resp.List[i] = ", resp.List[i])
		resp.List = append(resp.List, new(play.Result))
		resp.List[i].Id = sch.Id
		resp.List[i].PlayName = p.Name
		resp.List[i].Area = p.Area
		resp.List[i].Rating = p.Rating
		resp.List[i].Duration = p.Duration
		resp.List[i].ShowTime = sch.ShowTime
		resp.List[i].Price = p.Price
		resp.List[i].StudioName = resp1.Result.Name
		log.Println("Result = ", resp.List)
	}
	log.Println("Result = ", resp.List)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func PlayToScheduleService(ctx context.Context, req *play.PlayToScheduleRequest) (resp *play.PlayToScheduleResponse, err error) {
	resp = &play.PlayToScheduleResponse{BaseResp: &play.BaseResp{}}
	resp.Play, resp.ScheduleList, err = dao.PlayToSchedule(ctx, req.Id)
	log.Println("schedules = ", resp.ScheduleList)
	log.Println("play = ", resp.Play)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func GetScheduleService(ctx context.Context, req *play.GetScheduleRequest) (resp *play.GetScheduleResponse, err error) {
	log.Println(req)
	resp = &play.GetScheduleResponse{BaseResp: &play.BaseResp{}}
	resp.Schedule, err = dao.GetSchedule(ctx, req.Id)
	log.Println("schedule = ", resp.Schedule)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
