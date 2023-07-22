package models

import "github.com/APIVenda/db"

func Update(code int64, appscomercial Appscomercial) (int64, error) {
	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {
		return 0, err //Se algo der erro eu retorno
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	res, err := conn.Exec(`UPDATE appscomercial SET access_token=$1, client_id=$2, client_secret=$3 WHERE code=$4`, appscomercial.Access_token, appscomercial.Client_id, appscomercial.Client_secret)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
