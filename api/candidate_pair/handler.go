package candidate_pair

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartchoice/domain/candidate_pair"
	"strconv"
)

type Handler interface {
	GetCandidatePairByID(c *gin.Context)
}

type HandlerImpl struct {
	candidatePairService candidate_pair.Service
}

func NewCandidatePairHandler(candidatePairService candidate_pair.Service) *HandlerImpl {
	return &HandlerImpl{candidatePairService}
}

func (h *HandlerImpl) GetCandidatePairByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	candidatePair, err := h.candidatePairService.GetCandidatePairByID(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": candidatePair})
}
