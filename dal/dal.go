package dal

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

const DBUSER = "root"
const DBPASS = "123456"

var DB *sql.DB
var cfg mysql.Config

func InitDb() error {
	cfg = mysql.Config{
		User:   DBUSER,
		Passwd: DBPASS,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	DB = db

	return err
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
