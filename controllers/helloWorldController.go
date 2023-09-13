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
	g.JSON(http.StatusOK, gin.H{
		"success":               true,
		"message":               "We are are alive ):",
		"base_url":              "https://taofeeq-hng-stage-two.onrender.com/",
		"github_repo":           "https://github.com/harmlessprince/person_crud",
		"github_documentation":  "https://github.com/harmlessprince/person_crud#readme",
		"swagger_documentation": "https://taofeeq-hng-stage-two.onrender.com/docs/index.html",
	})
}
