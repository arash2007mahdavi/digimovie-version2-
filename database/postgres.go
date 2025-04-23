package database

import (
	logging "digimovie/src/Logging"
	"digimovie/src/config"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logging.NewLogger()
var DBClient *gorm.DB

func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Dbname, cfg.Postgres.Sslmode,
	)
	var err error

	DBClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := DBClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(*cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Info(logging.Postgres, logging.Startup, "postgres started", nil)
	return nil
}
