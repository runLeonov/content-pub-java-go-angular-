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

func (h *Handler) beAuthor(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	err = h.services.Account.ChangeRole("PRE_AUTHOR", id)
	c.JSON(http.StatusOK, err)
}

func (h *Handler) beUser(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	err = h.services.Account.ChangeRole("USER", id)
	c.JSON(http.StatusOK, err)
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

func (h *Handler) getPublishedTitles(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	titles, err := h.services.Account.GetUserPublishedByUserId(id)

	c.JSON(http.StatusOK, titles)
}

func (h *Handler) checkSubscribe(c *gin.Context) {
	authorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	author, err := h.services.Account.GetUserInfo(authorId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Account.GetUserInfo(id)
	sub, err := h.services.Account.CheckUserSubForAuthor(uint(author.ID), user)

	c.JSON(http.StatusOK, sub)
}

func (h *Handler) subscribe(c *gin.Context) {
	authorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	author, err := h.services.Account.GetUserInfo(authorId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Account.GetUserInfo(id)

	subscription := app.Subscription{
		Author:     author,
		Subscriber: user,
	}

	uid, err := h.services.Account.CreateSubscription(subscription)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": uid,
	})
}

func (h *Handler) getSubscriptions(c *gin.Context) {
	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.ParseToken(tok)
	authors, err := h.services.Account.GetUserSubscriptions(uint(id))
	c.JSON(http.StatusOK, authors)
}

func (h *Handler) unsubscribe(c *gin.Context) {
	authorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	tok, err := c.Cookie("jwt")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.ParseToken(tok)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	author, err := h.services.Account.GetUserInfo(authorId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Account.GetUserInfo(id)

	subscription := app.Subscription{
		Author:     author,
		Subscriber: user,
	}

	uid, err := h.services.Account.DeleteSubscription(subscription, user)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": uid,
	})
}

func (h *Handler) getCommentedTitlesForUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}
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
