package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proj1/domain/storage"
	"proj1/models"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func validation(c *gin.Context, name string, email string, phone string) (bool, string) {
	if strings.TrimSpace(name) == "" {
		return true, "Name cannot be left blank"
	}
	if !isEmailValid(email) {
		return true, "Invalid Email"
	}
	if len(phone) != 10 {
		return true, "Please Enter 10 Digit Phone Number"
	}
	return false, ""
}

func respondWithJSON(w http.ResponseWriter, reqcode int, eh interface{}) error {
	response, err := json.Marshal(eh)
	if err != nil {
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(reqcode)
	w.Write(response)
	return nil
}
func respondWithError(w http.ResponseWriter, code int, msg string) error {
	return respondWithJSON(w, code, map[string]string{"error": msg})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		respondWithError(c.Writer, 400, "Bad Request")
		return
	}
	flag, msg := validation(c, student.Name, student.Email, student.Phone)
	if flag {
		respondWithError(c.Writer, 400, msg)
		return
	}

	st, err := storage.Student.InsertStudent(student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Created Successfully": st})

}

func GetStudents(c *gin.Context) {
	student, err := storage.Student.GetStudent()
	if err != nil {
		respondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, student)

}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	st, err := storage.Student.GetStudentByID(id)
	if err != nil {
		respondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, st)

}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := c.BindJSON(&student)
	if err != nil {
		respondWithError(c.Writer, 400, "Bad Request")
		return
	}
	flag, msg := validation(c, student.Name, student.Email, student.Phone)
	if flag {
		respondWithError(c.Writer, 400, msg)
		return

	}
	_, err = storage.Student.UpdateStudent(student, id)
	if err != nil {
		respondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Updated Successfully"})
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	_, err := storage.Student.DeleteStudent(student, id)
	if err != nil {
		respondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully"})
}
