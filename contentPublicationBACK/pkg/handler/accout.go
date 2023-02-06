package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUserInfo(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	user, err := h.services.Account.GetUserInfo(id)

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getLikedTitles(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	titles, err := h.services.Account.GetLikedTitlesByUserId(id)

	c.JSON(http.StatusOK, titles)
}
