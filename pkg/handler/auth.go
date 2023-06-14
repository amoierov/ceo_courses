package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seo_courses/pkg/dto"
)

func (h *Handler) index(c *gin.Context) {
	c.File("C:\\Users\\kasat\\OneDrive\\Рабочий стол\\seo_courses\\cmd\\index.html")
}

func (h *Handler) admin(c *gin.Context) {
	c.File("C:\\Users\\kasat\\OneDrive\\Рабочий стол\\seo_courses\\cmd\\admin.html")
}

func (h *Handler) signUp(c *gin.Context) {
	var input dto.User

	if err := c.BindJSON(&input); err != nil { //BindJSON - метод парсинга тела json. Принимает ссылку на стуктуру, куда нужно парсить
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTopicsByCourse(c *gin.Context) {

	c.String(http.StatusOK, "Hello wrold")
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil { //BindJSON - метод парсинга тела json. Принимает ссылку на стуктуру, куда нужно парсить
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
