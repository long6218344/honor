package gorm

import (
	"database/sql"
	"time"

	"github.com/zhoushuguang/honor/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN string

	// database/sql
	MaxIdelConns int
	MaxOpenConns int
	MaxIdleTime  time.Duration
	MaxLifetime  time.Duration
}

type ORM struct {
	*gorm.DB
}

func NewORM(cfg *Config) *ORM {
	var (
		err   error
		db    *gorm.DB
		sqlDB *sql.DB
	)
	if db, err = gorm.Open(mysql.Open(cfg.DSN)); err != nil {
		log.Errorf("gorm.Open dsn: %s error: %v", cfg.DSN, err)
		panic(err)
	}
	if sqlDB, err = db.DB(); err != nil {
		log.Errorf("get database/sql db error: %v", err)
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(cfg.MaxIdleTime)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdelConns)

	return &ORM{db}
}

func (o *ORM) Close() {
	o.Close()
}
