package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcelofranco/api-tests-examples/contract-test/provider/data"
)

func (app *Config) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID!"})
		return
	}
	student, err := app.Repo.GetStudent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found!"})
		return
	}
	c.JSON(http.StatusOK, student)
}

type CreateStudentInput struct {
	Name string `json:"name" binding:"required"`
	CPF  string `json:"cpf" binding:"required"`
	RG   string `json:"rg" binding:"required"`
}

func (app *Config) CreateStudent(c *gin.Context) {
	var input CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var student = data.Student{
		Name: input.Name,
		CPF:  input.CPF,
		RG:   input.RG,
	}
	_, err := app.Repo.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating student!",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

type Class struct {
	ID         uint   `json:"id"`
	Discipline string `json:"discipline"`
	Day        string `json:"day"`
	Hour       string `json:"hour"`
}

func (app *Config) GetStudentClasses(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid student ID!",
		})
	}

	classes, err := app.Client.GetStudentClasses(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Class not found for student!",
		})
	}

	c.JSON(http.StatusOK, classes)
}
