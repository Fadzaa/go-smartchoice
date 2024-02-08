package debate

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartchoice/domain/debate"
)

type Handler interface {
	GetAllDebate(c *gin.Context)
}

type HandlerImpl struct {
	serviceDebate debate.Service
}

func NewDebateHandler(serviceDebate debate.Service) *HandlerImpl {
	return &HandlerImpl{serviceDebate}
}

func (h *HandlerImpl) GetAllDebate(c *gin.Context) {
	debates, err := h.serviceDebate.GetAllDebate()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": debates,
	})
}
