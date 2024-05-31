package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func GetMySQLConnection() string {
	mysqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))

	return mysqlConnection
}

func Connect() {
	var err error
	mysqlConnection := GetMySQLConnection()

	DB, err = sqlx.Connect("mysql", mysqlConnection)

	if err != nil {
		fmt.Println("bad connection")

		panic(err)
	}

	err = PingDB()
	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(2)

}

func PingDB() error {
	err := DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
