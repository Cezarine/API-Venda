package models

import "github.com/APIVenda/db"

func Insert(appscomercial Appscomercial) (code int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO appscomercial (access_token, client_id, client_secret) VALUES ($1, $2, $3) RETURNING code`
	err = conn.QueryRow(sql, appscomercial.Access_token, appscomercial.Client_id, appscomercial.Client_secret).Scan(&code)

	return
}
