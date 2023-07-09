package models

import "github.com/APIVenda/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {
		return 0, err //Se algo der erro eu retorno
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Done, todo.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
