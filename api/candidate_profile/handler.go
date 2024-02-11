package candidate_profile

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartchoice/domain/candidate_profile"
	"strconv"
)

type Handler interface {
	GetCandidateProfilebyID(c *gin.Context)
}

type HandlerImpl struct {
	candidateProfileService candidate_profile.Service
}

func NewCandidateProfileHandler(candidateProfileService candidate_profile.Service) *HandlerImpl {
	return &HandlerImpl{
		candidateProfileService: candidateProfileService,
	}
}

func (h *HandlerImpl) GetCandidateProfilebyID(c *gin.Context) {
	id := c.Param("id")
	//convert into uint
	candidateID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}

	candidateProfile, err := h.candidateProfileService.FindByID(candidateID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": candidateProfile})

}
