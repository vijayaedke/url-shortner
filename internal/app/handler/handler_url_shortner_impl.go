package handler

import (
	"net/http"

	"url-shortner/internal/app/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) URLShortner(c *gin.Context) {
	var request model.URLRequestResponse
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.service.URLShortner(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Redirect(c *gin.Context) {
	request := c.Query("url")
	response, err := h.service.Redirect(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
