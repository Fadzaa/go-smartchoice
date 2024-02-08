package party

import "github.com/gin-gonic/gin"

func RoutesParty(partyHandler Handler, router *gin.Engine) {
	p := router.Group("/party")
	p.GET("/", partyHandler.SearchAllParty)
}
