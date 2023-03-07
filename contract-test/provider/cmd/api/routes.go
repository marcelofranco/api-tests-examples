package main

import (
	"github.com/gin-gonic/gin"
)

const port = ":80"

func HandleRequests() {
	r := gin.Default()
	r.GET("/classes/:student", GetStudentClasses)
	r.Run(port)
}
