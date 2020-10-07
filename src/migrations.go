package src

import (
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/deliver-api-go/gen"
	"gopkg.in/gormigrate.v1"
)

// GetMigrations gets migrations
func GetMigrations(db *gen.DB) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		&gormigrate.Migration{
			ID: "0000_INIT",
			Migrate: func(tx *gorm.DB) error {
				return db.AutoMigrate()
			},
			Rollback: func(tx *gorm.DB) error {
				// there's not much we can do if initialization/automigration failed
				return nil
			},
		},
	}
}
