package controllers

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context)  {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func SearchStudentById(c *gin.Context)  {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func SearchStudentByCpf(c *gin.Context)  {
	cpf := c.Param("cpf")
	var student models.Student

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

// InsertStudent
//
//	@Summary		Add a new student
//	@Description	Router for add a new student
//	@ID				insert-student
//	@Accept			json
//	@Produce		json
//	@Param			student	body		models.Student	true	"Student model"
//	@Success		200		{object}	models.Student	"ok"
//	@Failure		400		{object}	http.StatusBadRequest
//	@Router			/students [post]
func InsertStudent(c *gin.Context)  {
	student := models.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context)  {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context)  {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.Delete(&student, id)
	c.JSON(http.StatusNoContent, gin.H{
		"data": "Student removed with success",
	})
}

func ShowIndexPage(c *gin.Context)  {
	var students []models.Student
	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}
