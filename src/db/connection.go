package db

import (
	"database/sql"
	"fmt"

	"github.com/APIVenda/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDB()

	stringConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.User, conf.Pass, conf.DataBase, conf.Port)

	conn, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err) //Nunca é bom usar panic em produção!!!
	}

	err = conn.Ping()
	return conn, err
}
