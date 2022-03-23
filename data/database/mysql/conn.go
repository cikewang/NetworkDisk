package mysql

import (
	cfg "NetworkDisk/data/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	db *sql.DB
)

func init() {
	db, _ = sql.Open("mysql", cfg.User+":"+cfg.Pass+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.DBName+"?charset="+cfg.Charset)
	db.SetMaxOpenConns(100)
	err := db.Ping()
	if err != nil {
		log.Println("123 Failed to connect to mysql, err : %s\n"+err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}
