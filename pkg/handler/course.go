package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seo_courses/pkg/dto"
	"strconv"
)

func (h *Handler) createCourse(c *gin.Context) {
	/*userId, err := getUserId(c)
	if err != nil {
		return
	}*/
	var input dto.Course
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Course.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []dto.Course `json:"data"`
}

func (h *Handler) getAllCourses(c *gin.Context) {
	/*userId, err := getUserId(c)
	if err != nil {
		return
	}*/

	courses, err := h.services.Course.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: courses,
	})
}

func (h *Handler) subscribeUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input dto.Subscribe
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Course.Subscribe(userId, input.CourseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) getCoursesByIdUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	/*id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}*/

	course, err := h.services.Course.GetCoursesByIdUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, course)

}

func (h *Handler) updateCourse(c *gin.Context) {

}

func (h *Handler) deleteCourse(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Course.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
