package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aungsannphyo/ywartalk/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBInstance *sql.DB
	once       sync.Once
)

func InitDb(dbCfg *config.MariadbConfig) *sql.DB {
	once.Do(func() {
		var err error

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbCfg.DBName))
		if err != nil {
			log.Fatalf("Failed to create database %s: %v", dbCfg.DBName, err)
		}

		dsn = buildDSN(dbCfg)
		DBInstance, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	})

	return DBInstance
}

func buildDSN(cfg *config.MariadbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}
