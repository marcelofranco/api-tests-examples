package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcelofranco/api-tests-examples/contract-test/provider/data"
)

func (app *Client) ListStudents(c *gin.Context) {
	var students []data.Student
	result := app.DB.Find(&students).Error
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Students not found",
		})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (app *Client) GetStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student data.Student
	if err := app.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found!"})
		return
	}
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found!",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

type Class struct {
	ID         uint   `json:"id"`
	Discipline string `json:"discipline"`
	Day        string `json:"day"`
	Hour       string `json:"hour"`
}

func (app *Client) GetStudentClasses(c *gin.Context) {
	id := c.Params.ByName("id")

	classesUrl := fmt.Sprintf("http://provider/classes/%s", id)

	request, err := http.NewRequest("GET", classesUrl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating student classes request",
		})
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error getting student classes",
		})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error getting student classes",
		})
		return
	}

	var classes []Class

	err = json.NewDecoder(response.Body).Decode(&classes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, classes)
}

type CreateStudentInput struct {
	Name string `json:"name" binding:"required"`
	CPF  string `json:"cpf" binding:"required"`
	RG   string `json:"rg" binding:"required"`
}

func (app *Client) CreateStudent(c *gin.Context) {
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
	if err := app.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating student!",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (app *Client) DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student data.Student
	if err := app.DB.Delete(&student, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting student!",
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
