package dao

import (
	"fmt"
	"reflect"

	"admin/proto"

	"gorm.io/gorm"
)

func parsePager(db *gorm.DB, pager *proto.Pager) *gorm.DB {
	if pager == nil {
		return db
	}

	return db.
		Offset(int(pager.Size * (pager.Number - 1))).
		Limit(int(pager.Size))
}

func parseSorter(db *gorm.DB, seqSorters []*proto.Sorter) *gorm.DB {
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

func parseConditionGroup(db *gorm.DB, condGroup *proto.ConditionGroup, model any) *gorm.DB {
	if condGroup == nil {
		return db
	}

	for _, cond := range condGroup.Conditions {
		if cond.GetField() == "" {
			fields := GetModelFields(model)
			nestedQuery := db
			for _, field := range fields {
				exp := fmt.Sprintf("%s %s ?", field, cond.GetOperator())
				nestedQuery = nestedQuery.Or(exp, cond.GetValue())
			}
			if condGroup.Concator == proto.Concator_AND {
				db = db.Where(nestedQuery)
			} else {
				db = db.Or(nestedQuery)
			}
		} else {
			exp := fmt.Sprintf("%s %s ?", cond.GetField(), cond.GetOperator())
			if condGroup.Concator == proto.Concator_AND {
				db = db.Where(exp, cond.GetValue())
			} else {
				db = db.Or(exp, cond.GetValue())
			}
		}
	}

	for _, nestGroup := range condGroup.ConditionGroups {
		nestedDB := parseConditionGroup(db, nestGroup, model)
		if condGroup.Concator == proto.Concator_AND {
			db = db.Where(nestedDB)
		} else {
			db = db.Or(nestedDB)
		}
	}

	return db
}

func GetModelFields(model any) []string {
	var fields []string
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	for i := 0; i < modelType.NumField(); i++ {
		fields = append(fields, modelType.Field(i).Name)
	}
	return fields
}
