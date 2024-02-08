package debate

import "github.com/gin-gonic/gin"

func RoutesDebate(debateHandler Handler, router *gin.Engine) {
	d := router.Group("/debate")
	d.GET("/", debateHandler.GetAllDebate)
}
