package di

import (
	"ecomb-go-demo/web/cfg"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"time"
)

func OpenDB(cfg cfg.MySqlConfig) (*gorm.DB, error) {
	dbLogger := logger.New(
		&log.Logger,
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful: false,
		},
	)

	conn, err := gorm.Open(mysql.Open(cfg.BuildDNS()), &gorm.Config{Logger: dbLogger})

	if err != nil {
		log.Error().Err(err).Msgf("cannot open DB %s", cfg.BuildDNS())
		return nil, err
	}

	db, _ := conn.DB()
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return conn, nil
}
