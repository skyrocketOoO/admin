package dao

import (
	"fmt"
	"reflect"

	"admin/internal/service/orm"
	"admin/proto"

	"github.com/skyrocketOoO/gorm-enhance-plugin/query"
	"gorm.io/gorm"
)

func applyPager(db *gorm.DB, pager *proto.Pager) *gorm.DB {
	if pager == nil {
		return db
	}

	return db.
		Offset(int(pager.Size * (pager.Number - 1))).
		Limit(int(pager.Size))
}

func applySorter(db *gorm.DB, seqSorters []*proto.Sorter) *gorm.DB {
	if len(seqSorters) == 0 {
		return db.Order("CreatedAt DESC")
	}
	for _, sorter := range seqSorters {
		expr := sorter.Field
		if !sorter.Ascending {
			expr += " desc"
		}
		db = db.Order(expr)
	}

	return db
}

func applyConditionGroup(
	db *gorm.DB,
	condGroup *proto.ConditionGroup,
) *gorm.DB {
	if condGroup == nil {
		return db
	}

	for _, cond := range condGroup.GetConditions() {
		exp := query.Build(cond.GetField(), cond.GetOperator())
		if condGroup.Concator == proto.Concator_AND {
			db = db.Where(exp, cond.GetValue())
		} else {
			db = db.Or(exp, cond.GetValue())
		}
	}

	// For nested groups
	blankDb := orm.GetDb()
	for _, nestGroup := range condGroup.GetConditionGroups() {
		nestedDB := blankDb.Session(&gorm.Session{NewDB: true})
		nestedDB = applyConditionGroup(nestedDB, nestGroup)

		if condGroup.Concator == proto.Concator_AND {
			db = db.Where(nestedDB)
		} else {
			db = db.Or(nestedDB)
		}
	}

	return db
}

// TODO: Not support for embedded and nested
func GetModelFields(model any) ([]string, error) {
	var fields []string
	modelType := reflect.TypeOf(model)

	// Check if the model is a slice
	if modelType.Kind() == reflect.Slice {
		// Get the element type of the slice (e.g., Model in []Model)
		modelType = modelType.Elem()
	}

	// Check if it's a pointer to a struct
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	// Ensure we're working with a struct type
	if modelType.Kind() != reflect.Struct {
		return nil, fmt.Errorf(
			"model must be a struct or a slice of structs, got %v",
			modelType.Kind(),
		)
	}

	// Extract field names from the struct type
	for i := 0; i < modelType.NumField(); i++ {
		if modelType.Field(i).Anonymous || modelType.Kind() == reflect.Struct {
			continue
		}
		fields = append(fields, modelType.Field(i).Name)
	}
	return fields, nil
}
