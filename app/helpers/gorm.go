package helpers

import (
	"fmt"

	"gorm.io/gorm"
)

func EqualCondition(db *gorm.DB, values map[string]interface{}) *gorm.DB {
	for field, value := range values {
		if value == nil {
			db = db.Where(fmt.Sprintf("%s IS NULL", field))
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}

	return db
}

func Preloads(db *gorm.DB, preloads ...string) *gorm.DB {
	for _, p := range preloads {
		db = db.Preload(p)
	}

	return db
}
