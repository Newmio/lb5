package handler

import (
	"fmt"
	"lb5/pkg/models"
	studentquery "lb5/pkg/postgres/studentQuery"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostStudent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := studentquery.Post(h.DB, student, userId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, models.Response{Response: fmt.Sprintf("student created, id: %d", id)})
}

func (h *Handler) GetAll(c *gin.Context) {
	if value := c.Query("param"); value == "1" {
		h.getAllStudentsJoinScore(c)
	} else {
		h.getAllStudents(c)
	}
}

func (h *Handler) getAllStudents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	students, err := studentquery.GetAll(h.DB, userId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}

func (h *Handler) getAllStudentsJoinScore(c *gin.Context) {
	students, err := studentquery.GetAllJoinScore(h.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}

func (h *Handler) GetByIdStudent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	student, err := studentquery.GetById(h.DB, userId, studentId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	var input models.Student
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = studentquery.Update(*h.DB, userId, studentId, input)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, input)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = studentquery.Delete(*h.DB, userId, studentId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{Response: "deleted"})
}

func (h *Handler) GroupByIdStudent(c *gin.Context) {
	students, err := studentquery.GroupById(h.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, students)
}

func (h *Handler) OrderByStudent(c *gin.Context) {
	students, err := studentquery.OrderBy(h.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}

type MinMax struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (h *Handler) GetByAgeStudent(c *gin.Context) {
	var input MinMax
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	students, err := studentquery.GetByAge(h.DB, input.Min, input.Max)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}
