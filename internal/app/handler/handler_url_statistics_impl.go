package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetURLStats(c *gin.Context) {
	response, err := h.service.GetURLStats(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
