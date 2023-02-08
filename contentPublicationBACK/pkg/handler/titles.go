package handler

import (
	app "contentPublicationBACK"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createTitle(c *gin.Context) {
	var input app.Title

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Titles.CreateTitle(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteTitle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	idt, err := h.services.Titles.DeleteTitle(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": idt,
	})
}

func (h *Handler) getTitles(c *gin.Context) {
	titles, err := h.services.Titles.GetTitles()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, titles)
}

func (h *Handler) getAllPossibleContent(c *gin.Context) {
	allContent, err := h.services.Titles.GetAllPossibleContent()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allContent)
}

func (h *Handler) getRandom(c *gin.Context) {
	title, err := h.services.Titles.GetRandom()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, title)
}

func (h *Handler) getTitlesByCategories(c *gin.Context) {
	var input []app.Category

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	titles, err := h.services.Titles.GetTitlesByCategories(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, titles)
}

func (h *Handler) getTitle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	title, err := h.services.Titles.GetTitle(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, title)
}

func (h *Handler) likeTitle(c *gin.Context) {
	var like app.Like
	if err := c.BindJSON(&like); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Titles.LikeOrUnlike(like, true)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, err)
}

func (h *Handler) updateViewsCount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}
	err = h.services.Titles.UpdateViewsForTitle(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) filterTitles(c *gin.Context) {
	var filter app.AllContent
	if err := c.BindJSON(&filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	titles, err := h.services.Titles.GetFilteredTitles(filter.Categories, filter.Tags, filter.Serials)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, titles)
}

func (h *Handler) unlikeTitle(c *gin.Context) {
	var like app.Like
	if err := c.BindJSON(&like); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Titles.LikeOrUnlike(like, false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, err)
}

func (h *Handler) updateTitle(c *gin.Context) {
	var input app.Title

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Titles.UpdateTitle(input, idP)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
