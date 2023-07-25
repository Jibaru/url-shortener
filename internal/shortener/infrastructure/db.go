package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"url-shortener/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(config *config.DBConfig) *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return db
}
