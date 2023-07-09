package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/APIVenda/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var todo models.Todo

	err := json.NewDecoder(c.Request.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"MEssage": fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
		}
	}

	c.JSON(http.StatusOK, resp)
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Printf("Erro ao buscar id na URL: %v", err)
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(c.Request.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}
	resp := map[string]any{
		"Error":   true,
		"Message": "dados atualizados com suscesso!",
	}
	c.JSON(http.StatusOK, resp)
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Printf("Erro ao buscar id na URL: %v", err)
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}
	resp := map[string]any{
		"Error":   true,
		"Message": "registro removido com suscesso!",
	}
	c.JSON(http.StatusOK, resp)
}

func List(c *gin.Context) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao buscar todos os registro: %v", err)
	}
	c.JSON(http.StatusOK, todos)
}

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Printf("Erro ao buscar id na URL: %v", err)
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	todo, err := models.GetTodo(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar o registro no banco: %v", err)
	}
	c.JSON(http.StatusOK, todo)
}
