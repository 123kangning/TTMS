package dao

import (
	"TTMS/kitex_gen/studio"
	"TTMS/kitex_gen/ticket"
	"context"
	"log"
)

func BatchAddTicket(ctx context.Context, ScheduleId int64, Price int32, PlayName string, StudioId int64, seats []*studio.Seat) error {
	SeatNum := len(seats)
	tickets := make([]*ticket.Ticket, SeatNum)
	for i, s := range seats {
		tickets[i] = new(ticket.Ticket)
		tickets[i].Price = Price
		tickets[i].PlayName = PlayName
		tickets[i].ScheduleId = ScheduleId
		tickets[i].SeatRow = int32(s.Row)
		tickets[i].SeatCol = int32(s.Col)
		tickets[i].StudioId = StudioId
	}
	return DB.WithContext(ctx).Create(&tickets).Error
}

func UpdateTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32, Price int32, status int32) error {
	t := ticket.Ticket{ScheduleId: ScheduleId, SeatRow: SeatRow, SeatCol: SeatCol, Price: Price, Status: status}
	return DB.WithContext(ctx).Where("schedule_id = ? and seat_row = ? and seat_col = ?", ScheduleId, SeatRow, SeatCol).Updates(&t).Error
}

func GetAllTicket(ctx context.Context, ScheduleId int64) ([]*ticket.Ticket, error) {
	var tickets []*ticket.Ticket
	err := DB.WithContext(ctx).Where("schedule_id = ?", ScheduleId).Find(&tickets).Error
	return tickets, err
}

func GetTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) *ticket.Ticket {
	t := ticket.Ticket{}
	DB.WithContext(ctx).Model(&ticket.Ticket{}).Where("schedule_id = ? and seat_row = ? and seat_col = ?", ScheduleId, SeatRow, SeatCol).Find(&t).Limit(1)
	return &t
}

func BuyTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) {
	err := DB.WithContext(ctx).Model(&ticket.Ticket{}).Where("schedule_id = ? and seat_row = ? and seat_col = ?", ScheduleId, SeatRow, SeatCol).Limit(1).Update("status", 1).Error
	if err != nil {
		log.Panicln(err)
	}
}

func ReturnTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) {
	//t := ticket.Ticket{}
	//DB.WithContext(ctx).Where("schedule_id = ? and seat_row = ? and seat_col = ?", ScheduleId, SeatRow, SeatCol).Find(&t)
	//if t.Status != 0 {
	err := DB.WithContext(ctx).Model(&ticket.Ticket{}).Where("schedule_id = ? and seat_row = ? and seat_col = ?", ScheduleId, SeatRow, SeatCol).Update("status", 0).Error
	if err != nil {
		log.Panicln(err)
	}
	//}
}
