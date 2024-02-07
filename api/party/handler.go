package party

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartchoice/domain/party"
)

type Handler interface {
	GetAllParty(c *gin.Context)
	SearchAllParty(c *gin.Context)
}

type HandlerImpl struct {
	serviceHandler party.Service
}

func NewPartyHandler(serviceParty party.Service) *HandlerImpl {
	return &HandlerImpl{serviceParty}
}

func (h *HandlerImpl) GetAllParty(c *gin.Context) {
	parties, err := h.serviceHandler.GetAll()
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

func (h *HandlerImpl) SearchAllParty(c *gin.Context) {
	query := c.Query("search")
	parties, err := h.serviceHandler.Search(query)
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
