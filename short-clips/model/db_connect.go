package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Migrator interface {
	Migrate(*gorm.DB)
}

func ConnectToDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}
	var err error
	dsn := "host=localhost user=shortclips password=shortclips dbname=shortclips"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %s", err.Error())
	}
	return DB, nil
}

func MigrateModels(db *gorm.DB, models []Migrator) {
	for _, m := range models {
		m.Migrate(db)
	}
}
