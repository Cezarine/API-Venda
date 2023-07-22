package main

import (
	"log"

	"github.com/APIVenda/config"
	handler "github.com/APIVenda/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.POST("/auth", handler.Create)
	r.PUT("/auth/:id", handler.Update)
	r.DELETE("/auth/:id", handler.Delete)
	r.GET("/auth", handler.List)
	r.GET("/auth/:id", handler.Get)
	r.GET("/", handler.Index)
	//r.NoRoute()
	r.Run(config.GetServerPort())

	//http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
