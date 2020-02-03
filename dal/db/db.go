package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func Init(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("fail to connect  mysql")
		return err
	}

	DB.SetMaxIdleConns(16)
	DB.SetMaxOpenConns(100)
	return nil
}
