package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySql interface{
	Close() error
	Query(query string, args ...string) (*sql.Rows, error)
	QueryRow(query string, args ...string) *sql.Row
	Exec(query string, args ...string) (sql.Result, error)
}

type mySql struct {
	DB *sql.DB
}

type MySqlConfig struct {
	DBName string
	Path string
}

func NewMySql(config MySqlConfig) (MySql, error) {
	db, err := sql.Open(config.DBName, config.Path)
	if err != nil {
		return nil, fmt.Errorf("mysql - couldn't open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("mysql - connection is not healthy: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return mySql{
		DB: db,
	},nil
}

func (m mySql) Close() error {
	return m.DB.Close()
}

func (m mySql) Query(query string, args ...string) (*sql.Rows, error) {
	return m.DB.Query(query, args)
}

func (m mySql) QueryRow(query string, args ...string) *sql.Row {
	return m.DB.QueryRow(query, args)
}

func (m mySql) Exec(query string, args ...string) (sql.Result, error) {
	return m.DB.Exec(query,args)
}
