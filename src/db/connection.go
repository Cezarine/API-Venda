package db

import (
	"database/sql"
	"fmt"

	"github.com/APIVenda/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDB()
	/* 							//POSTGRES
	stringConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.User, conf.Pass, conf.DataBase, conf.Port)
	conn, err := sql.Open("postgres", stringConnection)
	*/
	// MYSQL
	stringConnection := fmt.Sprintf("%s:%s@/%s", conf.User, conf.Pass, conf.DataBase)
	conn, err := sql.Open("mysql", stringConnection)
	if err != nil {
		panic(err) //Nunca é bom usar panic em produção!!!
	}

	err = conn.Ping()
	return conn, err
}
