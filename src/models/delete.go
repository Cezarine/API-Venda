package models

import "github.com/APIVenda/db"

func Delete(client_id int64) (int64, error) {
	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {
		return 0, err //Se algo der erro eu retorno
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	res, err := conn.Exec(`DELETE FROM appscomercial WHERE client_id=$1`, client_id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
