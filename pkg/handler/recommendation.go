package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seo_courses/pkg/dto"
	"strconv"
)

func (h *Handler) GetForm(c *gin.Context) {
	formId, err := strconv.Atoi(c.Param("form_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	questions, err := h.services.Form.GetForm(formId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, questions)
}

func (h *Handler) UserPreferences(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input dto.CreateAnswerRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Form.UserPreferences(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getRecommedationResponse struct {
	Data []dto.RecommendationRequest `json:"data"`
}

func (h *Handler) createRecommendation(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	recom, err := h.services.Form.CreateRecommendation(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, recom)
}
