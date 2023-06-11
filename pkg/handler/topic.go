package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seo_courses/pkg/dto"
)

func (h *Handler) createTopic(c *gin.Context) {
	courseId, err := getCourseId(c)
	if err != nil {
		return
	}

	var input dto.Topic
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Topic.Create(courseId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllTopics(c *gin.Context) {

}

func (h *Handler) getTopicById(c *gin.Context) {

}

func (h *Handler) updateTopic(c *gin.Context) {

}

func (h *Handler) deleteTopic(c *gin.Context) {

}
