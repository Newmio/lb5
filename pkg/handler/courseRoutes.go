package handler

import (
	"fmt"
	"lb5/pkg/models"
	coursequery "lb5/pkg/postgres/courseQuery"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostCourse(c *gin.Context) {
	var course models.Course
	if err := c.BindJSON(&course); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := coursequery.Post(h.DB, studentId, course)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, models.Response{Response: fmt.Sprintf("course created, id: %d", id)})
}

func (h *Handler) GetAllCourse(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	courses, err := coursequery.GetAll(h.DB, studentId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, courses)
}

func (h *Handler) GetByIdCourse(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	courseId, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	course, err := coursequery.GetById(h.DB, studentId, courseId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, course)
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	var input models.Course
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	courseId, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = coursequery.Update(h.DB, studentId, courseId, input)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, input)
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	courseId, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = coursequery.Delete(h.DB, studentId, courseId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{Response: "deleted"})
}
