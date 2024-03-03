package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/ars0915/gogolook-exercise/config"
)

func NewDB(config config.ConfENV) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.SQLite.Database), &gorm.Config{})
}
