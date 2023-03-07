package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudentClasses(c *gin.Context) {
	student, _ := strconv.Atoi(c.Params.ByName("student"))
	var classes []Class

	result := DB.Model(&Class{}).Joins("inner join student_classes as cs on cs.class_id = classes.id", DB.Where(&StudentClass{StudentId: uint(student)})).Scan(&classes).Error

	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Class not found for student",
		})
		return
	}

	c.JSON(http.StatusOK, classes)
}
