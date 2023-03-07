package main

import (
	"github.com/gin-gonic/gin"
)

func HandleRequests(port string) {
	r := gin.Default()
	r.GET("/classes/:student", GetStudentClasses)
	r.Run(port)
}
