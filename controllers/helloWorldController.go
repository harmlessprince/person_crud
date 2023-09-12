package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api

// HelloWorld PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} hello world
// @Router /api [get]

func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "hello world")
}
