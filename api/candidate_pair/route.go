package candidate_pair

import "github.com/gin-gonic/gin"

func RouteCandidatePair(candidatePairHandler Handler, router *gin.Engine) {
	c := router.Group("/candidate_pair")
	c.GET("/:id", candidatePairHandler.GetCandidatePairByID)
}
