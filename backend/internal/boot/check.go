package boot

import (
	"errors"
	"reflect"

	"admin/internal/global"
	"admin/internal/model"
	"admin/internal/service/orm"
)

func Check() (err error) {
	db := orm.GetDb()

	masterRole := model.Role{}
	if err := db.
		Where("ID = ?", global.MasterRoleID).
		Take(&masterRole).Error; err != nil {
		return err
	}
	if !allBoolsTrue(masterRole.Page) {
		return errors.New("master role must has all permissions")
	}

	defaultRole := model.Role{}
	if err := db.
		Where("ID = ?", global.DefaultRoleID).
		Take(&defaultRole).Error; err != nil {
		return err
	}
	if !allBoolsFalse(defaultRole.Page) {
		return errors.New("default role must has no permissions")
	}

	return
}

// Checks if all boolean fields are true
func allBoolsTrue(v interface{}) bool {
	val := reflect.ValueOf(v)
	return checkAllBools(val, true)
}

// Checks if all boolean fields are false
func allBoolsFalse(v interface{}) bool {
	val := reflect.ValueOf(v)
	return checkAllBools(val, false)
}

// Helper function to check all bools for a specific target (true or false)
func checkAllBools(val reflect.Value, target bool) bool {
	switch val.Kind() {
	case reflect.Ptr:
		// Dereference pointers
		if val.IsNil() {
			return target // Consider nil pointer as target for this check
		}
		return checkAllBools(val.Elem(), target)
	case reflect.Struct:
		// Iterate over struct fields
		for i := 0; i < val.NumField(); i++ {
			if !checkAllBools(val.Field(i), target) {
				return false
			}
		}
	case reflect.Bool:
		// Check if the boolean value matches the target (true or false)
		return val.Bool() == target
	}
	// For other kinds (e.g., int, string, etc.), we don't care
	return true
}
