package handler

import (
	app "contentPublicationBACK"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) getUserInfoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}
	user, err := h.services.Account.GetUserInfo(id)
	c.JSON(http.StatusOK, user)
}

func (h *Handler) changeUserInfo(c *gin.Context) {
	var inputUser app.User

	if err := c.BindJSON(&inputUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.ParseToken(tok)
	user, err := h.services.Account.ChangeUserInfo(inputUser, id)

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

func (h *Handler) getLikedTitlesLimit(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	title, err := h.services.Account.GetLikedTitlesByUserIdLimit(id)

	c.JSON(http.StatusOK, title)
}

func (h *Handler) getCommentedTitles(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	titles, err := h.services.Account.GetCommentedTitlesByUserId(id)

	c.JSON(http.StatusOK, titles)
}

func (h *Handler) getCommentedTitlesLimit(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	title, err := h.services.Account.GetCommentedTitlesByUserIdLimit(id)

	c.JSON(http.StatusOK, title)
}
