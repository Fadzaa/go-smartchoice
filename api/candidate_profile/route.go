package candidate_profile

import "github.com/gin-gonic/gin"

func RoutesCandidateProfile(handler Handler, router *gin.Engine) {
	cp := router.Group("/candidate_profile")
	cp.GET("/:id", handler.GetCandidateProfilebyID)
}
