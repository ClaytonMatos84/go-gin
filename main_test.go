package main

import (
	"api-go-gin/controllers"
	"api-go-gin/database"
	"api-go-gin/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var CPF = "98745679988"

func SetupRouterTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func CreateStudentOnDb()  {
	student := models.Student{Name: "Aluno mock", CPF: CPF, RG: "123654989"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentOnDb()  {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifiedStatusCodeOnSalutationWithParameter(t *testing.T)  {
	r := SetupRouterTest()
	r.GET("/salutation/:name", controllers.Salutation)
	req, _ := http.NewRequest("GET", "/salutation/marco", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	// if response.Code != http.StatusOK {
	// 	t.Fatalf("Status error: value %d - expected %d", response.Code, http.StatusOK)
	// }
	mockResponse := `{"Seja bem-vindo":"marco. Tudo bom?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, mockResponse, string(responseBody))
}

func TestShowStudents(t *testing.T)  {
	database.ConnectDB()
	CreateStudentOnDb()
	defer DeleteStudentOnDb()

	r := SetupRouterTest()
	r.GET("/students", controllers.ShowStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentById(t *testing.T)  {
	database.ConnectDB()
	CreateStudentOnDb()
	defer DeleteStudentOnDb()

	r := SetupRouterTest()
	r.GET("/students/:id", controllers.SearchStudentById)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var mockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, int(ID), int(mockStudent.ID))
}

func TestSearchStudentByCpf(t *testing.T)  {
	database.ConnectDB()
	CreateStudentOnDb()
	defer DeleteStudentOnDb()

	r := SetupRouterTest()
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCpf)
	path := "/students/cpf/" + CPF
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var mockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, CPF, mockStudent.CPF)
}

func TestUpdateStudent(t *testing.T)  {
	database.ConnectDB()
	CreateStudentOnDb()
	defer DeleteStudentOnDb()

	r := SetupRouterTest()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	updateStudent := models.Student{Name: "Nome editado", RG: "335559998", CPF: "12345678947"}
	jsonValue, _ := json.Marshal(updateStudent)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var mockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, updateStudent.Name, mockStudent.Name)
	assert.Equal(t, updateStudent.RG, mockStudent.RG)
	assert.Equal(t, updateStudent.CPF, mockStudent.CPF)
}

func TestDeleteStudent(t *testing.T)  {
	database.ConnectDB()
	CreateStudentOnDb()

	r := SetupRouterTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNoContent, response.Code)
}
