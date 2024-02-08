package infrastructure

import (
	"github.com/gin-gonic/gin"
	"smartchoice/api/debate"
	"smartchoice/api/party"
)

func SetupRoutes(partyHandler party.Handler, debateHandler debate.Handler) *gin.Engine {
	router := gin.Default()

	party.RoutesParty(partyHandler, router)
	debate.RoutesDebate(debateHandler, router)

	return router
}
