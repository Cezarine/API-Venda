package models

import (
	"github.com/APIVenda/db"
)

func GetAppscomercial(client_id int64) (appscomercial Appscomercial, err error) {
	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {
		return //Se algo der erro eu retorno, como no cabeçalho da função eu já disse que o retorno será (appscomercial e err), técnicamente eu já declarei as variaveis, então só preciso colocar "return"
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	sql := `SELECT * FROM todos WHERE client_id = $1`                                                                        //Crio o SELECT
	row := conn.QueryRow(sql, client_id)                                                                                     //Chamo a função para executar o comando
	err = row.Scan(&appscomercial.Code, &appscomercial.Access_token, &appscomercial.Client_id, &appscomercial.Client_secret) //E faço um Scan para verficar se não deu nenhum error
	return
}

func GetAll() (appscomercials []Appscomercial, err error) {

	conn, err := db.OpenConnection() //Conecto com o Banco de Dados
	if err != nil {                  //Se algo der erro eu retorno sem nada para tratar isso no HANDLER
		return
	}
	defer conn.Close() //Mato a Conexão se nada der errado

	sql := `SELECT * FROM appscomercial` //Crio o SELECT
	rows, err := conn.Query(sql)         //Chamo a função para executar o comando
	if err != nil {
		return
	}
	for rows.Next() {
		var appscomercial Appscomercial
		err = rows.Scan(&appscomercial.Code, &appscomercial.Access_token, &appscomercial.Client_id, &appscomercial.Client_secret) //E faço um Scan para verficar se não deu nenhum error
		if err != nil {
			continue
		}
		appscomercials = append(appscomercials, appscomercial) //O "appscomercials" com S no final é o "appscomercial" retornado da função
	}
	return
}
