package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQl struct {
	db *sqlx.DB
	tx *sqlx.Tx
	// showSql bool
	// log     Logger
}

type Config struct {
	Enable       bool   `toml:"enable" json:"enable"`
	Driver       string `toml:"driver" json:"driver"`
	Dsn          string `toml:"dsn" json:"dsn"`
	MaxOpenConns int    `toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns" json:"max_idle_conns"`
	MaxLifetime  int    `toml:"max_lefttime" json:"max_lefttime"`
	MaxIdletime  int    `toml:"max_idle_time" json:"max_idle_time"`
	ShowSql      bool   `toml:"show_sql" json:"show_sql"`
}

func NewMySQL(configValue *Config) *MySQl {
	fmt.Println(configValue)
	db, err := sqlx.Connect(configValue.Driver, configValue.Dsn)
	if err != nil {
		fmt.Println("MySQL connection error: ", err)
		panic(err)
	}

	db.SetMaxIdleConns(configValue.MaxIdleConns)
	db.SetMaxOpenConns(configValue.MaxOpenConns)
	if configValue.MaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(configValue.MaxLifetime) * time.Second)
	}
	if configValue.MaxIdletime > 0 {
		db.SetConnMaxIdleTime(time.Duration(configValue.MaxIdletime) * time.Second)
	}

	return &MySQl{
		db: db,
		// showSql: configValue.ShowSql,
		// log:     &DefaultLogger{},
	}
}
