package coreapi

import (
	"net/http"

	"github.com/djengua/raffle-api/core"
	"github.com/gin-gonic/gin"
)

func (s *Server) homePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title":   "Home Page",
		"message": "Hello World",
	})
}

func (s *Server) TicketSuggest(ctx *gin.Context) {

	result3 := []string{}
	result6 := []string{}
	result9 := []string{}

	for i := 0; i < 3; i++ {
		result3 = append(result3, core.TicketNumber(5))
	}
	for i := 0; i < 6; i++ {
		result6 = append(result6, core.TicketNumber(5))
	}
	for i := 0; i < 9; i++ {
		result9 = append(result9, core.TicketNumber(5))
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data": map[string]interface{}{
			"result3": result3,
			"result6": result6,
			"result9": result9,
		},
	})
}

func (s *Server) MelSuggest(ctx *gin.Context) {

	result := core.MelGenerator(6)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  result,
	})
}
