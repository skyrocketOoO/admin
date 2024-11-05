package dao

import (
	"fmt"

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

func parseConditionGroup(db *gorm.DB, condGroup *proto.ConditionGroup) *gorm.DB {
	if condGroup == nil {
		return db
	}

	for _, cond := range condGroup.Conditions {
		exp := fmt.Sprintf("%s %s ?", cond.GetField(), cond.GetOperator())
		if condGroup.Concator == proto.Concator_AND {
			db = db.Where(exp, cond.GetValue())
		} else {
			db = db.Or(exp, cond.GetValue())
		}
	}

	for _, nestGroup := range condGroup.ConditionGroups {
		nestedDB := parseConditionGroup(db, nestGroup)
		if condGroup.Concator == proto.Concator_AND {
			db = db.Where(nestedDB)
		} else {
			db = db.Or(nestedDB)
		}
	}

	return db
}
