package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddBatchSeat(ctx context.Context, studioInfo *studio.Studio) error {
	id, row, col := int(studioInfo.Id), int(studioInfo.RowsCount), int(studioInfo.ColsCount)
	seats := make([]*studio.Seat, 0, row*col)
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			seats = append(seats, &studio.Seat{StudioId: int64(id), Row: int64(i), Col: int64(j), Status: 1})
		}
	}
	return DB.WithContext(ctx).Create(seats).Error
}
func AddSeat(ctx context.Context, seatInfo *studio.Seat) error {
	s := studio.Seat{}
	if DB.WithContext(ctx).Where("studio_id = ? and row = ? and col = ?", seatInfo.StudioId, seatInfo.Row, seatInfo.Col).Find(&s); s.Id > 0 && s.Status > 0 {
		fmt.Println("不允许的行为：同一位置重复加入座位")
		return errors.New("不允许的行为：同一位置重复加入座位")
	} else if s.Id == 0 {
		//没有该座位的数据存在
		return DB.WithContext(ctx).Create(seatInfo).Error
	} else {
		//该座位的数据存在，只是status==0
		return UpdateSeat(ctx, seatInfo)
	}

}
func GetAllSeat(ctx context.Context, studioId, Current, PageSize int) ([]*studio.Seat, error) {
	seats := make([]*studio.Seat, PageSize)
	tx := DB.WithContext(ctx).Where("studio_id = ?", studioId).Order("row").Order("col").Offset((Current - 1) * PageSize).Limit(PageSize).Find(&seats)
	return seats, tx.Error
}
func UpdateSeat(ctx context.Context, seatInfo *studio.Seat) error {
	s := studio.Seat{}
	if DB.WithContext(ctx).Where("studio_id = ? and row = ? and col = ? ", seatInfo.StudioId, seatInfo.Row, seatInfo.Col).Find(&s); s.Id > 0 {
		return DB.WithContext(ctx).Model(&s).Where("studio_id = ? and row = ? and col = ? ", seatInfo.StudioId, seatInfo.Row, seatInfo.Col).Update("status", seatInfo.Status).Error
	}
	if seatInfo.Status == 0 { //删除了一个座位
		DB.WithContext(ctx).Model(&studio.Studio{}).Where("id = ?", seatInfo.StudioId).UpdateColumn("seats_count", gorm.Expr("seats_count - 1"))
	}
	return errors.New("该位置上无座位，修改失败")
}
func DeleteSeat(ctx context.Context, seatInfo *studio.Seat) error {
	seatInfo.Status = 0
	return UpdateSeat(ctx, seatInfo)
}
