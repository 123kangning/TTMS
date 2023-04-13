// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package order

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMessage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Order) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Order[number], err)
}

func (x *Order) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Order) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Order) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Order) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Order) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Order) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Date, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Order) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.Value, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Order) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddOrderRequest[number], err)
}

func (x *AddOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AddOrderRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AddOrderRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddOrderRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddOrderRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AddOrderRequest) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Value, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddOrderResponse[number], err)
}

func (x *AddOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *CommitOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CommitOrderRequest[number], err)
}

func (x *CommitOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CommitOrderRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CommitOrderRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CommitOrderRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CommitOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CommitOrderResponse[number], err)
}

func (x *CommitOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *DeleteOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteOrderRequest[number], err)
}

func (x *DeleteOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteOrderResponse[number], err)
}

func (x *DeleteOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *UpdateOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateOrderRequest[number], err)
}

func (x *UpdateOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateOrderResponse[number], err)
}

func (x *UpdateOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetAllOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetAllOrderRequest[number], err)
}

func (x *GetAllOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetAllOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetAllOrderResponse[number], err)
}

func (x *GetAllOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetAllOrderResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Order
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.List = append(x.List, &v)
	return offset, nil
}

func (x *OrderAnalysis) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderAnalysis[number], err)
}

func (x *OrderAnalysis) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PlayId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PlayName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayArea, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PlayDuration, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.StartData, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.EndData, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.TotalTicket, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Sales, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *OrderAnalysis) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetOrderAnalysisRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetOrderAnalysisRequest[number], err)
}

func (x *GetOrderAnalysisRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PlayId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetOrderAnalysisResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetOrderAnalysisResponse[number], err)
}

func (x *GetOrderAnalysisResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetOrderAnalysisResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v OrderAnalysis
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.OrderAnalysis = &v
	return offset, nil
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMessage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMessage)
	return offset
}

func (x *Order) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Order) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Order) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *Order) fastWriteField3(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.ScheduleId)
	return offset
}

func (x *Order) fastWriteField4(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.SeatRow)
	return offset
}

func (x *Order) fastWriteField5(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.SeatCol)
	return offset
}

func (x *Order) fastWriteField6(buf []byte) (offset int) {
	if x.Date == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.Date)
	return offset
}

func (x *Order) fastWriteField7(buf []byte) (offset int) {
	if x.Value == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 7, x.Value)
	return offset
}

func (x *Order) fastWriteField8(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 8, x.Type)
	return offset
}

func (x *AddOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *AddOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *AddOrderRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ScheduleId)
	return offset
}

func (x *AddOrderRequest) fastWriteField3(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.SeatRow)
	return offset
}

func (x *AddOrderRequest) fastWriteField4(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.SeatCol)
	return offset
}

func (x *AddOrderRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Data == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.Data)
	return offset
}

func (x *AddOrderRequest) fastWriteField6(buf []byte) (offset int) {
	if x.Value == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.Value)
	return offset
}

func (x *AddOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AddOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *CommitOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *CommitOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *CommitOrderRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ScheduleId)
	return offset
}

func (x *CommitOrderRequest) fastWriteField3(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.SeatRow)
	return offset
}

func (x *CommitOrderRequest) fastWriteField4(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.SeatCol)
	return offset
}

func (x *CommitOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CommitOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *DeleteOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *DeleteOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *UpdateOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *UpdateOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetAllOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetAllOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *GetAllOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetAllOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetAllOrderResponse) fastWriteField2(buf []byte) (offset int) {
	if x.List == nil {
		return offset
	}
	for i := range x.List {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.List[i])
	}
	return offset
}

func (x *OrderAnalysis) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	return offset
}

func (x *OrderAnalysis) fastWriteField1(buf []byte) (offset int) {
	if x.PlayId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.PlayId)
	return offset
}

func (x *OrderAnalysis) fastWriteField2(buf []byte) (offset int) {
	if x.PlayName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.PlayName)
	return offset
}

func (x *OrderAnalysis) fastWriteField3(buf []byte) (offset int) {
	if x.PlayArea == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.PlayArea)
	return offset
}

func (x *OrderAnalysis) fastWriteField4(buf []byte) (offset int) {
	if x.PlayDuration == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.PlayDuration)
	return offset
}

func (x *OrderAnalysis) fastWriteField5(buf []byte) (offset int) {
	if x.StartData == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.StartData)
	return offset
}

func (x *OrderAnalysis) fastWriteField6(buf []byte) (offset int) {
	if x.EndData == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.EndData)
	return offset
}

func (x *OrderAnalysis) fastWriteField7(buf []byte) (offset int) {
	if x.TotalTicket == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.TotalTicket)
	return offset
}

func (x *OrderAnalysis) fastWriteField8(buf []byte) (offset int) {
	if x.Sales == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.Sales)
	return offset
}

func (x *OrderAnalysis) fastWriteField9(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 9, x.Price)
	return offset
}

func (x *GetOrderAnalysisRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetOrderAnalysisRequest) fastWriteField1(buf []byte) (offset int) {
	if x.PlayId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.PlayId)
	return offset
}

func (x *GetOrderAnalysisResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetOrderAnalysisResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetOrderAnalysisResponse) fastWriteField2(buf []byte) (offset int) {
	if x.OrderAnalysis == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.OrderAnalysis)
	return offset
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.StatusCode)
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.StatusMessage == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMessage)
	return n
}

func (x *Order) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Order) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Order) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *Order) sizeField3() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.ScheduleId)
	return n
}

func (x *Order) sizeField4() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.SeatRow)
	return n
}

func (x *Order) sizeField5() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.SeatCol)
	return n
}

func (x *Order) sizeField6() (n int) {
	if x.Date == "" {
		return n
	}
	n += fastpb.SizeString(6, x.Date)
	return n
}

func (x *Order) sizeField7() (n int) {
	if x.Value == 0 {
		return n
	}
	n += fastpb.SizeInt32(7, x.Value)
	return n
}

func (x *Order) sizeField8() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(8, x.Type)
	return n
}

func (x *AddOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *AddOrderRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *AddOrderRequest) sizeField2() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ScheduleId)
	return n
}

func (x *AddOrderRequest) sizeField3() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.SeatRow)
	return n
}

func (x *AddOrderRequest) sizeField4() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.SeatCol)
	return n
}

func (x *AddOrderRequest) sizeField5() (n int) {
	if x.Data == "" {
		return n
	}
	n += fastpb.SizeString(5, x.Data)
	return n
}

func (x *AddOrderRequest) sizeField6() (n int) {
	if x.Value == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.Value)
	return n
}

func (x *AddOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AddOrderResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *CommitOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *CommitOrderRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *CommitOrderRequest) sizeField2() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ScheduleId)
	return n
}

func (x *CommitOrderRequest) sizeField3() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.SeatRow)
	return n
}

func (x *CommitOrderRequest) sizeField4() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.SeatCol)
	return n
}

func (x *CommitOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CommitOrderResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *DeleteOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteOrderRequest) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *DeleteOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteOrderResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *UpdateOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateOrderRequest) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *UpdateOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateOrderResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetAllOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetAllOrderRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *GetAllOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetAllOrderResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetAllOrderResponse) sizeField2() (n int) {
	if x.List == nil {
		return n
	}
	for i := range x.List {
		n += fastpb.SizeMessage(2, x.List[i])
	}
	return n
}

func (x *OrderAnalysis) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	return n
}

func (x *OrderAnalysis) sizeField1() (n int) {
	if x.PlayId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.PlayId)
	return n
}

func (x *OrderAnalysis) sizeField2() (n int) {
	if x.PlayName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.PlayName)
	return n
}

func (x *OrderAnalysis) sizeField3() (n int) {
	if x.PlayArea == "" {
		return n
	}
	n += fastpb.SizeString(3, x.PlayArea)
	return n
}

func (x *OrderAnalysis) sizeField4() (n int) {
	if x.PlayDuration == "" {
		return n
	}
	n += fastpb.SizeString(4, x.PlayDuration)
	return n
}

func (x *OrderAnalysis) sizeField5() (n int) {
	if x.StartData == "" {
		return n
	}
	n += fastpb.SizeString(5, x.StartData)
	return n
}

func (x *OrderAnalysis) sizeField6() (n int) {
	if x.EndData == "" {
		return n
	}
	n += fastpb.SizeString(6, x.EndData)
	return n
}

func (x *OrderAnalysis) sizeField7() (n int) {
	if x.TotalTicket == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.TotalTicket)
	return n
}

func (x *OrderAnalysis) sizeField8() (n int) {
	if x.Sales == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.Sales)
	return n
}

func (x *OrderAnalysis) sizeField9() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeInt32(9, x.Price)
	return n
}

func (x *GetOrderAnalysisRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetOrderAnalysisRequest) sizeField1() (n int) {
	if x.PlayId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.PlayId)
	return n
}

func (x *GetOrderAnalysisResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetOrderAnalysisResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetOrderAnalysisResponse) sizeField2() (n int) {
	if x.OrderAnalysis == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.OrderAnalysis)
	return n
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMessage",
}

var fieldIDToName_Order = map[int32]string{
	1: "Id",
	2: "UserId",
	3: "ScheduleId",
	4: "SeatRow",
	5: "SeatCol",
	6: "Date",
	7: "Value",
	8: "Type",
}

var fieldIDToName_AddOrderRequest = map[int32]string{
	1: "UserId",
	2: "ScheduleId",
	3: "SeatRow",
	4: "SeatCol",
	5: "Data",
	6: "Value",
}

var fieldIDToName_AddOrderResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_CommitOrderRequest = map[int32]string{
	1: "UserId",
	2: "ScheduleId",
	3: "SeatRow",
	4: "SeatCol",
}

var fieldIDToName_CommitOrderResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_DeleteOrderRequest = map[int32]string{
	1: "Id",
}

var fieldIDToName_DeleteOrderResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_UpdateOrderRequest = map[int32]string{
	1: "Id",
}

var fieldIDToName_UpdateOrderResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_GetAllOrderRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetAllOrderResponse = map[int32]string{
	1: "BaseResp",
	2: "List",
}

var fieldIDToName_OrderAnalysis = map[int32]string{
	1: "PlayId",
	2: "PlayName",
	3: "PlayArea",
	4: "PlayDuration",
	5: "StartData",
	6: "EndData",
	7: "TotalTicket",
	8: "Sales",
	9: "Price",
}

var fieldIDToName_GetOrderAnalysisRequest = map[int32]string{
	1: "PlayId",
}

var fieldIDToName_GetOrderAnalysisResponse = map[int32]string{
	1: "BaseResp",
	2: "OrderAnalysis",
}
