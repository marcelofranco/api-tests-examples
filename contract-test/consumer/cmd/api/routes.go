package main

import (
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	r := gin.Default()
	r.GET("/students/:id", app.GetStudent)
	r.GET("/students/:id/classes", app.GetStudentClasses)
	r.POST("/students", app.CreateStudent)
	return r
}
