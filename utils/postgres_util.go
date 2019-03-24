package utils

import (
	"database/sql"
	"fmt"
	// "github.com/go-xorm/core"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
    host     = "121.42.160.108"
    port     = 5432
    user     = "postgres"
    password = "tuji2013"
    dbname   = "torch_service"
)
type singleton struct{}
var ins *singleton
var once sync.Once

func GetIns() *singleton {
    once.Do(func(){
        ins = &singleton{}
    })
    return ins
}

func PostgresConnect() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=verify-full",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
       log.Println(err)
	}

	return db
}

