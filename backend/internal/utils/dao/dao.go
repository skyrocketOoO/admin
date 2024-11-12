package dao

import (
	"context"
	"time"

	"admin/proto"

	"gorm.io/gorm"
)

// Get page items with total count, src must pass the pointer of item
func ListWithPager(
	db *gorm.DB,
	option *proto.ListOption,
	src any,
) (total int64, err error) {
	db = applySorter(db, option.GetSorters())
	db = applyConditionGroup(db, option.GetConditionGroup())

	if err = db.Count(&total).Error; err != nil {
		return
	}

	db = applyPager(db, option.GetPager())

	err = db.Find(src).Error
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
