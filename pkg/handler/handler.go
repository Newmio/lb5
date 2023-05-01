package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB *sqlx.DB
}

func InitRoutes(db *sqlx.DB, router *gin.Engine) {
	h := &Handler{DB: db}

	auth := router.Group("/auth")
	{
		auth.POST("/sign", h.GetUser)
		auth.POST("register", h.PostUser)
	}

	api := router.Group("/api", h.userIdentity)
	{
		student := api.Group("/student")
		{
			student.GET("/", h.GetAll)
			student.GET("/:id", h.GetByIdStudent)
			student.POST("/", h.PostStudent)
			student.PUT("/:id", h.UpdateStudent)
			student.DELETE("/:id", h.DeleteStudent)

			course := student.Group(":id/course")
			{
				course.GET("/:course_id", h.GetByIdCourse)
				course.PUT("/:course_id", h.UpdateCourse)
				course.DELETE("/:course_id", h.DeleteCourse)
				course.GET("/", h.GetAllCourse)
				course.POST("/", h.PostCourse)

				score := course.Group(":course_id/score")
				{
					score.GET("/", h.GetAllScore)
					score.GET("/:score_id", h.GetByIdScore)
					score.POST("/", h.PostScore)
					score.PUT("/:score_id", h.UpdateScore)
					score.DELETE("/:score_id", h.DeleteScore)
				}
			}
		}
	}

	router.GET("/group_by", h.GroupByIdStudent)
	router.GET("/order_by", h.OrderByStudent)
	router.GET("/age", h.GetByAgeStudent)
}
