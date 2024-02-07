package infrastructure

import (
	"github.com/gin-gonic/gin"
	"smartchoice/api/party"
)

func SetupRoutes(partyHandler party.Handler) *gin.Engine {
	router := gin.Default()

	p := router.Group("/party")
	p.GET("/", partyHandler.GetAllParty)
	p.GET("/filter", partyHandler.SearchAllParty)

	return router
}
