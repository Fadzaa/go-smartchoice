package party

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartchoice/domain/party"
)

type Handler interface {
	SearchAllParty(c *gin.Context)
}

type HandlerImpl struct {
	serviceHandler party.Service
}

func NewPartyHandler(serviceParty party.Service) *HandlerImpl {
	return &HandlerImpl{serviceParty}
}

func (h *HandlerImpl) SearchAllParty(c *gin.Context) {
	query := c.Query("search")
	parties, err := h.serviceHandler.GetAllParty(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": parties,
	})
}
