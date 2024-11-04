package dao

import (
	"context"
	"time"

	"admin/proto"

	"gorm.io/gorm"
)

// Get page items with total count, src must pass the pointer of item
func ListWithPager(
	tx *gorm.DB,
	option *proto.ListOption,
	src any,
) (total int64, err error) {
	// Assemble order
	if len(option.Sorters) != 0 {
		for _, sorter := range option.Sorters {
			expr := sorter.Field
			if !sorter.Ascending {
				expr += " desc"
			}
			db = db.Order(expr)
		}
	} else {
		db = db.Order("CreatedAt DESC")
	}

	if option.FilterGroup != nil {
		for _, q := range option.Query {
			if q.Fuzzy {
				db = db.Where(q.Field+" LIKE ?", "%"+q.Value+"%")
			} else {
				db = db.Where(q.Field+" = ?", q.Value)
			}
		}
	}

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if option.Pager != (Pager{}) {
		db.Offset(int(option.Pager.Size * (option.Pager.Number - 1))).
			Limit(int(option.Pager.Size))
	}

	err = db.Find(src).Error
	return
}

func ParseConditionGroup(tx *gorm.DB, condGroup *proto.ConditionGroup) {
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
