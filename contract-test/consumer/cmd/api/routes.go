package main

import (
	"github.com/gin-gonic/gin"
)

func (app *Client) routes() *gin.Engine {
	r := gin.Default()
	r.GET("/students", app.ListStudents)
	r.GET("/students/:id", app.GetStudent)
	r.GET("/students/:id/classes", app.GetStudentClasses)
	r.POST("/students", app.CreateStudent)
	r.DELETE("/students/:id", app.DeleteStudent)
	return r
}
