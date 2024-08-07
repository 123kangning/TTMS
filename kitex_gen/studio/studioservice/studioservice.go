// Code generated by Kitex v0.9.1. DO NOT EDIT.

package studioservice

import (
	studio "TTMS/kitex_gen/studio"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AddStudio": kitex.NewMethodInfo(
		addStudioHandler,
		newAddStudioArgs,
		newAddStudioResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetAllStudio": kitex.NewMethodInfo(
		getAllStudioHandler,
		newGetAllStudioArgs,
		newGetAllStudioResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateStudio": kitex.NewMethodInfo(
		updateStudioHandler,
		newUpdateStudioArgs,
		newUpdateStudioResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteStudio": kitex.NewMethodInfo(
		deleteStudioHandler,
		newDeleteStudioArgs,
		newDeleteStudioResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AddSeat": kitex.NewMethodInfo(
		addSeatHandler,
		newAddSeatArgs,
		newAddSeatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetAllSeat": kitex.NewMethodInfo(
		getAllSeatHandler,
		newGetAllSeatArgs,
		newGetAllSeatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateSeat": kitex.NewMethodInfo(
		updateSeatHandler,
		newUpdateSeatArgs,
		newUpdateSeatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteSeat": kitex.NewMethodInfo(
		deleteSeatHandler,
		newDeleteSeatArgs,
		newDeleteSeatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetStudio": kitex.NewMethodInfo(
		getStudioHandler,
		newGetStudioArgs,
		newGetStudioResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	studioServiceServiceInfo                = NewServiceInfo()
	studioServiceServiceInfoForClient       = NewServiceInfoForClient()
	studioServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return studioServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return studioServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return studioServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "StudioService"
	handlerType := (*studio.StudioService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "studio",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func addStudioHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.AddStudioRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).AddStudio(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddStudioArgs:
		success, err := handler.(studio.StudioService).AddStudio(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddStudioResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddStudioArgs() interface{} {
	return &AddStudioArgs{}
}

func newAddStudioResult() interface{} {
	return &AddStudioResult{}
}

type AddStudioArgs struct {
	Req *studio.AddStudioRequest
}

func (p *AddStudioArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.AddStudioRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddStudioArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddStudioArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddStudioArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddStudioArgs) Unmarshal(in []byte) error {
	msg := new(studio.AddStudioRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddStudioArgs_Req_DEFAULT *studio.AddStudioRequest

func (p *AddStudioArgs) GetReq() *studio.AddStudioRequest {
	if !p.IsSetReq() {
		return AddStudioArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddStudioArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddStudioArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddStudioResult struct {
	Success *studio.AddStudioResponse
}

var AddStudioResult_Success_DEFAULT *studio.AddStudioResponse

func (p *AddStudioResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.AddStudioResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddStudioResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddStudioResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddStudioResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddStudioResult) Unmarshal(in []byte) error {
	msg := new(studio.AddStudioResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddStudioResult) GetSuccess() *studio.AddStudioResponse {
	if !p.IsSetSuccess() {
		return AddStudioResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddStudioResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.AddStudioResponse)
}

func (p *AddStudioResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddStudioResult) GetResult() interface{} {
	return p.Success
}

func getAllStudioHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.GetAllStudioRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).GetAllStudio(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetAllStudioArgs:
		success, err := handler.(studio.StudioService).GetAllStudio(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAllStudioResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetAllStudioArgs() interface{} {
	return &GetAllStudioArgs{}
}

func newGetAllStudioResult() interface{} {
	return &GetAllStudioResult{}
}

type GetAllStudioArgs struct {
	Req *studio.GetAllStudioRequest
}

func (p *GetAllStudioArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.GetAllStudioRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAllStudioArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAllStudioArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAllStudioArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetAllStudioArgs) Unmarshal(in []byte) error {
	msg := new(studio.GetAllStudioRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAllStudioArgs_Req_DEFAULT *studio.GetAllStudioRequest

func (p *GetAllStudioArgs) GetReq() *studio.GetAllStudioRequest {
	if !p.IsSetReq() {
		return GetAllStudioArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAllStudioArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetAllStudioArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetAllStudioResult struct {
	Success *studio.GetAllStudioResponse
}

var GetAllStudioResult_Success_DEFAULT *studio.GetAllStudioResponse

func (p *GetAllStudioResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.GetAllStudioResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAllStudioResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAllStudioResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAllStudioResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetAllStudioResult) Unmarshal(in []byte) error {
	msg := new(studio.GetAllStudioResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllStudioResult) GetSuccess() *studio.GetAllStudioResponse {
	if !p.IsSetSuccess() {
		return GetAllStudioResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAllStudioResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.GetAllStudioResponse)
}

func (p *GetAllStudioResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAllStudioResult) GetResult() interface{} {
	return p.Success
}

func updateStudioHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.UpdateStudioRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).UpdateStudio(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateStudioArgs:
		success, err := handler.(studio.StudioService).UpdateStudio(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateStudioResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateStudioArgs() interface{} {
	return &UpdateStudioArgs{}
}

func newUpdateStudioResult() interface{} {
	return &UpdateStudioResult{}
}

type UpdateStudioArgs struct {
	Req *studio.UpdateStudioRequest
}

func (p *UpdateStudioArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.UpdateStudioRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateStudioArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateStudioArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateStudioArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateStudioArgs) Unmarshal(in []byte) error {
	msg := new(studio.UpdateStudioRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateStudioArgs_Req_DEFAULT *studio.UpdateStudioRequest

func (p *UpdateStudioArgs) GetReq() *studio.UpdateStudioRequest {
	if !p.IsSetReq() {
		return UpdateStudioArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateStudioArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateStudioArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateStudioResult struct {
	Success *studio.UpdateStudioResponse
}

var UpdateStudioResult_Success_DEFAULT *studio.UpdateStudioResponse

func (p *UpdateStudioResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.UpdateStudioResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateStudioResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateStudioResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateStudioResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateStudioResult) Unmarshal(in []byte) error {
	msg := new(studio.UpdateStudioResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateStudioResult) GetSuccess() *studio.UpdateStudioResponse {
	if !p.IsSetSuccess() {
		return UpdateStudioResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateStudioResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.UpdateStudioResponse)
}

func (p *UpdateStudioResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateStudioResult) GetResult() interface{} {
	return p.Success
}

func deleteStudioHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.DeleteStudioRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).DeleteStudio(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteStudioArgs:
		success, err := handler.(studio.StudioService).DeleteStudio(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteStudioResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteStudioArgs() interface{} {
	return &DeleteStudioArgs{}
}

func newDeleteStudioResult() interface{} {
	return &DeleteStudioResult{}
}

type DeleteStudioArgs struct {
	Req *studio.DeleteStudioRequest
}

func (p *DeleteStudioArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.DeleteStudioRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteStudioArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteStudioArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteStudioArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteStudioArgs) Unmarshal(in []byte) error {
	msg := new(studio.DeleteStudioRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteStudioArgs_Req_DEFAULT *studio.DeleteStudioRequest

func (p *DeleteStudioArgs) GetReq() *studio.DeleteStudioRequest {
	if !p.IsSetReq() {
		return DeleteStudioArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteStudioArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteStudioArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteStudioResult struct {
	Success *studio.DeleteStudioResponse
}

var DeleteStudioResult_Success_DEFAULT *studio.DeleteStudioResponse

func (p *DeleteStudioResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.DeleteStudioResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteStudioResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteStudioResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteStudioResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteStudioResult) Unmarshal(in []byte) error {
	msg := new(studio.DeleteStudioResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteStudioResult) GetSuccess() *studio.DeleteStudioResponse {
	if !p.IsSetSuccess() {
		return DeleteStudioResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteStudioResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.DeleteStudioResponse)
}

func (p *DeleteStudioResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteStudioResult) GetResult() interface{} {
	return p.Success
}

func addSeatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.AddSeatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).AddSeat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddSeatArgs:
		success, err := handler.(studio.StudioService).AddSeat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddSeatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddSeatArgs() interface{} {
	return &AddSeatArgs{}
}

func newAddSeatResult() interface{} {
	return &AddSeatResult{}
}

type AddSeatArgs struct {
	Req *studio.AddSeatRequest
}

func (p *AddSeatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.AddSeatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddSeatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddSeatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddSeatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddSeatArgs) Unmarshal(in []byte) error {
	msg := new(studio.AddSeatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddSeatArgs_Req_DEFAULT *studio.AddSeatRequest

func (p *AddSeatArgs) GetReq() *studio.AddSeatRequest {
	if !p.IsSetReq() {
		return AddSeatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddSeatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddSeatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddSeatResult struct {
	Success *studio.AddSeatResponse
}

var AddSeatResult_Success_DEFAULT *studio.AddSeatResponse

func (p *AddSeatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.AddSeatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddSeatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddSeatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddSeatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddSeatResult) Unmarshal(in []byte) error {
	msg := new(studio.AddSeatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddSeatResult) GetSuccess() *studio.AddSeatResponse {
	if !p.IsSetSuccess() {
		return AddSeatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddSeatResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.AddSeatResponse)
}

func (p *AddSeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddSeatResult) GetResult() interface{} {
	return p.Success
}

func getAllSeatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.GetAllSeatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).GetAllSeat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetAllSeatArgs:
		success, err := handler.(studio.StudioService).GetAllSeat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAllSeatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetAllSeatArgs() interface{} {
	return &GetAllSeatArgs{}
}

func newGetAllSeatResult() interface{} {
	return &GetAllSeatResult{}
}

type GetAllSeatArgs struct {
	Req *studio.GetAllSeatRequest
}

func (p *GetAllSeatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.GetAllSeatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAllSeatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAllSeatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAllSeatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetAllSeatArgs) Unmarshal(in []byte) error {
	msg := new(studio.GetAllSeatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAllSeatArgs_Req_DEFAULT *studio.GetAllSeatRequest

func (p *GetAllSeatArgs) GetReq() *studio.GetAllSeatRequest {
	if !p.IsSetReq() {
		return GetAllSeatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAllSeatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetAllSeatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetAllSeatResult struct {
	Success *studio.GetAllSeatResponse
}

var GetAllSeatResult_Success_DEFAULT *studio.GetAllSeatResponse

func (p *GetAllSeatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.GetAllSeatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAllSeatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAllSeatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAllSeatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetAllSeatResult) Unmarshal(in []byte) error {
	msg := new(studio.GetAllSeatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllSeatResult) GetSuccess() *studio.GetAllSeatResponse {
	if !p.IsSetSuccess() {
		return GetAllSeatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAllSeatResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.GetAllSeatResponse)
}

func (p *GetAllSeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAllSeatResult) GetResult() interface{} {
	return p.Success
}

func updateSeatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.UpdateSeatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).UpdateSeat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateSeatArgs:
		success, err := handler.(studio.StudioService).UpdateSeat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateSeatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateSeatArgs() interface{} {
	return &UpdateSeatArgs{}
}

func newUpdateSeatResult() interface{} {
	return &UpdateSeatResult{}
}

type UpdateSeatArgs struct {
	Req *studio.UpdateSeatRequest
}

func (p *UpdateSeatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.UpdateSeatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateSeatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateSeatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateSeatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateSeatArgs) Unmarshal(in []byte) error {
	msg := new(studio.UpdateSeatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateSeatArgs_Req_DEFAULT *studio.UpdateSeatRequest

func (p *UpdateSeatArgs) GetReq() *studio.UpdateSeatRequest {
	if !p.IsSetReq() {
		return UpdateSeatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateSeatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateSeatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateSeatResult struct {
	Success *studio.UpdateSeatResponse
}

var UpdateSeatResult_Success_DEFAULT *studio.UpdateSeatResponse

func (p *UpdateSeatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.UpdateSeatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateSeatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateSeatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateSeatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateSeatResult) Unmarshal(in []byte) error {
	msg := new(studio.UpdateSeatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateSeatResult) GetSuccess() *studio.UpdateSeatResponse {
	if !p.IsSetSuccess() {
		return UpdateSeatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateSeatResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.UpdateSeatResponse)
}

func (p *UpdateSeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateSeatResult) GetResult() interface{} {
	return p.Success
}

func deleteSeatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.DeleteSeatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).DeleteSeat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteSeatArgs:
		success, err := handler.(studio.StudioService).DeleteSeat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteSeatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteSeatArgs() interface{} {
	return &DeleteSeatArgs{}
}

func newDeleteSeatResult() interface{} {
	return &DeleteSeatResult{}
}

type DeleteSeatArgs struct {
	Req *studio.DeleteSeatRequest
}

func (p *DeleteSeatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.DeleteSeatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteSeatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteSeatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteSeatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteSeatArgs) Unmarshal(in []byte) error {
	msg := new(studio.DeleteSeatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteSeatArgs_Req_DEFAULT *studio.DeleteSeatRequest

func (p *DeleteSeatArgs) GetReq() *studio.DeleteSeatRequest {
	if !p.IsSetReq() {
		return DeleteSeatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteSeatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteSeatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteSeatResult struct {
	Success *studio.DeleteSeatResponse
}

var DeleteSeatResult_Success_DEFAULT *studio.DeleteSeatResponse

func (p *DeleteSeatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.DeleteSeatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteSeatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteSeatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteSeatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteSeatResult) Unmarshal(in []byte) error {
	msg := new(studio.DeleteSeatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteSeatResult) GetSuccess() *studio.DeleteSeatResponse {
	if !p.IsSetSuccess() {
		return DeleteSeatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteSeatResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.DeleteSeatResponse)
}

func (p *DeleteSeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteSeatResult) GetResult() interface{} {
	return p.Success
}

func getStudioHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(studio.GetStudioRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(studio.StudioService).GetStudio(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetStudioArgs:
		success, err := handler.(studio.StudioService).GetStudio(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetStudioResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetStudioArgs() interface{} {
	return &GetStudioArgs{}
}

func newGetStudioResult() interface{} {
	return &GetStudioResult{}
}

type GetStudioArgs struct {
	Req *studio.GetStudioRequest
}

func (p *GetStudioArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(studio.GetStudioRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetStudioArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetStudioArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetStudioArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetStudioArgs) Unmarshal(in []byte) error {
	msg := new(studio.GetStudioRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetStudioArgs_Req_DEFAULT *studio.GetStudioRequest

func (p *GetStudioArgs) GetReq() *studio.GetStudioRequest {
	if !p.IsSetReq() {
		return GetStudioArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetStudioArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetStudioArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetStudioResult struct {
	Success *studio.GetStudioResponse
}

var GetStudioResult_Success_DEFAULT *studio.GetStudioResponse

func (p *GetStudioResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(studio.GetStudioResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetStudioResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetStudioResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetStudioResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetStudioResult) Unmarshal(in []byte) error {
	msg := new(studio.GetStudioResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetStudioResult) GetSuccess() *studio.GetStudioResponse {
	if !p.IsSetSuccess() {
		return GetStudioResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetStudioResult) SetSuccess(x interface{}) {
	p.Success = x.(*studio.GetStudioResponse)
}

func (p *GetStudioResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetStudioResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddStudio(ctx context.Context, Req *studio.AddStudioRequest) (r *studio.AddStudioResponse, err error) {
	var _args AddStudioArgs
	_args.Req = Req
	var _result AddStudioResult
	if err = p.c.Call(ctx, "AddStudio", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllStudio(ctx context.Context, Req *studio.GetAllStudioRequest) (r *studio.GetAllStudioResponse, err error) {
	var _args GetAllStudioArgs
	_args.Req = Req
	var _result GetAllStudioResult
	if err = p.c.Call(ctx, "GetAllStudio", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateStudio(ctx context.Context, Req *studio.UpdateStudioRequest) (r *studio.UpdateStudioResponse, err error) {
	var _args UpdateStudioArgs
	_args.Req = Req
	var _result UpdateStudioResult
	if err = p.c.Call(ctx, "UpdateStudio", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteStudio(ctx context.Context, Req *studio.DeleteStudioRequest) (r *studio.DeleteStudioResponse, err error) {
	var _args DeleteStudioArgs
	_args.Req = Req
	var _result DeleteStudioResult
	if err = p.c.Call(ctx, "DeleteStudio", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddSeat(ctx context.Context, Req *studio.AddSeatRequest) (r *studio.AddSeatResponse, err error) {
	var _args AddSeatArgs
	_args.Req = Req
	var _result AddSeatResult
	if err = p.c.Call(ctx, "AddSeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllSeat(ctx context.Context, Req *studio.GetAllSeatRequest) (r *studio.GetAllSeatResponse, err error) {
	var _args GetAllSeatArgs
	_args.Req = Req
	var _result GetAllSeatResult
	if err = p.c.Call(ctx, "GetAllSeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateSeat(ctx context.Context, Req *studio.UpdateSeatRequest) (r *studio.UpdateSeatResponse, err error) {
	var _args UpdateSeatArgs
	_args.Req = Req
	var _result UpdateSeatResult
	if err = p.c.Call(ctx, "UpdateSeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteSeat(ctx context.Context, Req *studio.DeleteSeatRequest) (r *studio.DeleteSeatResponse, err error) {
	var _args DeleteSeatArgs
	_args.Req = Req
	var _result DeleteSeatResult
	if err = p.c.Call(ctx, "DeleteSeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetStudio(ctx context.Context, Req *studio.GetStudioRequest) (r *studio.GetStudioResponse, err error) {
	var _args GetStudioArgs
	_args.Req = Req
	var _result GetStudioResult
	if err = p.c.Call(ctx, "GetStudio", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
