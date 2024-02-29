package ports

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/periclescesar/rinha-2024-q1-go/configs"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db == nil {
		db = connect()
	}

	return db
}

func connect() *sql.DB {
	dbConf := configs.Configs.DbConf

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Pass, dbConf.Dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return conn
}

func Disconnect() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
