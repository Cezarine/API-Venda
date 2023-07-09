package models

import (
	"github.com/APIVenda/db"
)

func GetTodo(id int64) (todo Todo, err error) {

	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {
		return //Se algo der erro eu retorno, como no cabeçalho da função eu já disse que o retorno será (todo e err), técnicamente eu já declarei as variaveis, então só preciso colocar "return"
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	sql := `SELECT * FROM todos WHERE id = $1`                           //Crio o SELECT
	row := conn.QueryRow(sql, id)                                        //Chamo a função para executar o comando
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done) //E faço um Scan para verficar se não deu nenhum error
	return
}

func GetAll() (todos []Todo, err error) {

	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {                  //Se algo der erro eu retorno sem nada para tratar isso no HANDLER
		return
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	sql := `SELECT * FROM todos` //Crio o SELECT
	rows, err := conn.Query(sql) //Chamo a função para executar o comando
	if err != nil {
		return
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done) //E faço um Scan para verficar se não deu nenhum error
		if err != nil {
			continue
		}
		todos = append(todos, todo) //O "todos" com S no final é o "todo" retornado da função
	}
	return
}
