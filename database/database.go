package database

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const DriverName = "mysql"

var cfg *mysql.Config

func Connect() *sqlx.DB {
	db, err := sqlx.Connect(DriverName, cfg.FormatDSN())
	if err != nil {
		panic("Could not connect to the database")
	}
	return db
}

func init() {
	c := mysql.NewConfig()
	c.Addr = os.Getenv("DBADDR")
	c.User = os.Getenv("DBUSER")
	c.Passwd = os.Getenv("DBPASS")
	c.DBName = os.Getenv("DBNAME")
	c.Net = os.Getenv("DBNET")
	cfg = c
}
