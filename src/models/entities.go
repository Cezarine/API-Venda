package models

type Todo struct { //Todo, significa "pêndencia", ou seja, esse MODEL é um pêndencia, tem que existir ele, ele é um principal
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"` //Done, significa "feito", ou seja, se deu certo ou não
}
