package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	Pager struct {
		Number int
		Size   int
	}
)

// Get page items with total count, src must pass the pointer of item
func ListWithPager(db *gorm.DB, pager Pager, src any) (total int64, err error) {
	if err = db.Count(&total).Error; err != nil {
		return
	}
	err = db.Offset(pager.Size * (pager.Number - 1)).Limit(pager.Size).Find(src).Error
	return
}

func Ping(db *gorm.DB, ctx context.Context) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := sqlDb.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
