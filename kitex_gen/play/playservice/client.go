// Code generated by Kitex v0.4.4. DO NOT EDIT.

package playservice

import (
	play "TTMS/kitex_gen/play"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddPlay(ctx context.Context, Req *play.AddPlayRequest, callOptions ...callopt.Option) (r *play.AddPlayResponse, err error)
	UpdatePlay(ctx context.Context, Req *play.UpdatePlayRequest, callOptions ...callopt.Option) (r *play.UpdatePlayResponse, err error)
	DeletePlay(ctx context.Context, Req *play.DeletePlayRequest, callOptions ...callopt.Option) (r *play.DeletePlayResponse, err error)
	GetAllPlay(ctx context.Context, Req *play.GetAllPlayRequest, callOptions ...callopt.Option) (r *play.GetAllPlayResponse, err error)
	AddSchedule(ctx context.Context, Req *play.AddScheduleRequest, callOptions ...callopt.Option) (r *play.AddScheduleResponse, err error)
	UpdateSchedule(ctx context.Context, Req *play.UpdateScheduleRequest, callOptions ...callopt.Option) (r *play.UpdateScheduleResponse, err error)
	DeleteSchedule(ctx context.Context, Req *play.DeleteScheduleRequest, callOptions ...callopt.Option) (r *play.DeleteScheduleResponse, err error)
	GetAllSchedule(ctx context.Context, Req *play.GetAllScheduleRequest, callOptions ...callopt.Option) (r *play.GetAllScheduleResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPlayServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPlayServiceClient struct {
	*kClient
}

func (p *kPlayServiceClient) AddPlay(ctx context.Context, Req *play.AddPlayRequest, callOptions ...callopt.Option) (r *play.AddPlayResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddPlay(ctx, Req)
}

func (p *kPlayServiceClient) UpdatePlay(ctx context.Context, Req *play.UpdatePlayRequest, callOptions ...callopt.Option) (r *play.UpdatePlayResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePlay(ctx, Req)
}

func (p *kPlayServiceClient) DeletePlay(ctx context.Context, Req *play.DeletePlayRequest, callOptions ...callopt.Option) (r *play.DeletePlayResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePlay(ctx, Req)
}

func (p *kPlayServiceClient) GetAllPlay(ctx context.Context, Req *play.GetAllPlayRequest, callOptions ...callopt.Option) (r *play.GetAllPlayResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllPlay(ctx, Req)
}

func (p *kPlayServiceClient) AddSchedule(ctx context.Context, Req *play.AddScheduleRequest, callOptions ...callopt.Option) (r *play.AddScheduleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddSchedule(ctx, Req)
}

func (p *kPlayServiceClient) UpdateSchedule(ctx context.Context, Req *play.UpdateScheduleRequest, callOptions ...callopt.Option) (r *play.UpdateScheduleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateSchedule(ctx, Req)
}

func (p *kPlayServiceClient) DeleteSchedule(ctx context.Context, Req *play.DeleteScheduleRequest, callOptions ...callopt.Option) (r *play.DeleteScheduleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSchedule(ctx, Req)
}

func (p *kPlayServiceClient) GetAllSchedule(ctx context.Context, Req *play.GetAllScheduleRequest, callOptions ...callopt.Option) (r *play.GetAllScheduleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllSchedule(ctx, Req)
}
