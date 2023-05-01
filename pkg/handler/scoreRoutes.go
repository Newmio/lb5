package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"lb5/pkg/models"
	scorequery "lb5/pkg/postgres/scoreQuery"
)

func (h *Handler) PostScore(c *gin.Context) {
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

	var input models.Score
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := scorequery.Post(h.DB, input, studentId, courseId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, fmt.Sprintf("%d", id))
}

func (h *Handler) GetAllScore(c *gin.Context) {
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

	scores, err := scorequery.GetAll(h.DB, studentId, courseId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, scores)
}

func (h *Handler) GetByIdScore(c *gin.Context) {
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
	scoreId, err := strconv.Atoi(c.Param("score_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	score, err := scorequery.GetById(h.DB, courseId, studentId, scoreId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, score)
}

func (h *Handler) UpdateScore(c *gin.Context) {
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
	scoreId, err := strconv.Atoi(c.Param("score_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	var input models.Score
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = scorequery.Update(h.DB, studentId, courseId, scoreId, input)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, input)
}

func (h *Handler) DeleteScore(c *gin.Context) {
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
	scoreId, err := strconv.Atoi(c.Param("score_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = scorequery.Delete(h.DB, courseId, studentId, scoreId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{Response: "deleted"})
}

func (h *Handler) GetByCourseScore(c *gin.Context) {
	var courseId int
	if err := c.BindJSON(&courseId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
}
