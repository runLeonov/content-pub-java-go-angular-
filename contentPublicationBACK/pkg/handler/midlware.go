package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "no token")
		return
	}
	userId, err := h.services.Authorization.ParseToken(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "id not found in params")
		return 0, errors.New("id noy found in params")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "cannot parse id")
		return 0, errors.New("cannot parse id")
	}
	return idInt, nil
}
