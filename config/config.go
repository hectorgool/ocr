package config

import (
	"ocr/schema"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
)

func init() {

	// Allocate a logical processor for every available core.
	runtime.GOMAXPROCS(runtime.NumCPU())

	db, err = sqlx.Connect("mysql", "santo:asdf@(db:3306)/bunsan")
	if err != nil {
		panic(err.Error())
	}
	GetDB().MustExec(schema.Schema)
}

//GetDB gets mysql conection
func GetDB() *sqlx.DB {
	return db
}
