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
	model any,
) (total int64, err error) {
	db = parseSorter(db, option.GetSorters())
	db = parseConditionGroup(db, option.GetConditionGroup(), model)

	if err = db.Count(&total).Error; err != nil {
		return
	}

	db = parsePager(db, option.GetPager())

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
