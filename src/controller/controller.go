package controller

import (
	"net/http"

	"github.com/APIVenda/models"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "", nil)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetUsuarios(c *gin.Context) {
	var usuarios []models.Appcomercial

	//db.DB.Find(&usuarios)
	c.HTML(http.StatusOK, "index.html", gin.H{"usuarios": usuarios})
}
