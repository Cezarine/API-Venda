package main

import (
	"fmt"
	"net/http"

	"github.com/APIVenda/config"
	handler "github.com/APIVenda/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/", handler.Create)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
	r.GET("/", handler.List)
	r.GET("/:id", handler.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
