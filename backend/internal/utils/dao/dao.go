package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	ListOption struct {
		Pager  Pager
		Sorter Sorter
		Query  []Query
	}

	Pager struct {
		Number int32
		Size   int32
	}

	Sorter struct {
		Field string
		Asc   bool
	}

	Query struct {
		Field string
		Value string
		Fuzzy bool
	}
)

// Get page items with total count, src must pass the pointer of item
func ListWithPager(db *gorm.DB, option ListOption, src any) (total int64, err error) {
	if option.Sorter != (Sorter{}) {
		expr := option.Sorter.Field
		if !option.Sorter.Asc {
			expr += " desc"
		}
		db = db.Order(expr)
	} else {
		db = db.Order("CreatedAt DESC")
	}

	if option.Query != nil {
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
		db.Offset(int(option.Pager.Size * (option.Pager.Number - 1))).Limit(int(option.Pager.Size))
	}

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
