package handlers

import (
	"fmt"
	"net/http"
	"proj1/api/utils"
	"proj1/domain/filters"
	"proj1/domain/pagination"
	"proj1/domain/storage"
	"proj1/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		utils.RespondWithError(c.Writer, 400, "Bad Request")
		return
	}

	student.Phone = strings.ReplaceAll(student.Phone, "-", "")

	flag, msg := utils.Validation(c, student.Name, student.Email, student.Phone)
	if flag {
		utils.RespondWithError(c.Writer, 400, msg)
		return
	}
	if CheckEmailExists(c, student.Email) {
		utils.RespondWithError(c.Writer, 400, "Email already Exists")
		return
	}
	if CheckPhoneExists(c, student.Phone) {
		utils.RespondWithError(c.Writer, 400, "Phone Number already Exists")
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

	sp, err := pagination.NewPaginationFromGinCtx(c)
	if err != nil {
		utils.RespondWithError(c.Writer, 500, err.Error())
		return
	}

	sf, err := filters.NewStudentFilterFromCtx(c)
	if err != nil {
		utils.RespondWithError(c.Writer, 500, err.Error())
		return
	}

	sts, err := storage.Student.GetStudents(sp, sf)
	if err != nil {
		utils.RespondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.NewDefaultResponse(sts, sp))
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	st, err := storage.Student.GetStudentByID(id)
	if err != nil {
		utils.RespondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, st)

}

func CheckEmailExists(c *gin.Context, email string) bool {
	_, err := storage.Student.CheckEmailExists(email)
	return err == nil
}

func CheckPhoneExists(c *gin.Context, phone string) bool {
	_, err := storage.Student.CheckPhoneExists(phone)
	return err == nil
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := c.BindJSON(&student)
	if err != nil {
		utils.RespondWithError(c.Writer, 400, "Bad Request")
		return
	}
	flag, msg := utils.Validation(c, student.Name, student.Email, student.Phone)
	if flag {
		utils.RespondWithError(c.Writer, 400, msg)
		return

	}
	_, err = storage.Student.UpdateStudent(student, id)
	if err != nil {
		utils.RespondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Updated Successfully"})
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	_, err := storage.Student.DeleteStudent(student, id)
	if err != nil {
		utils.RespondWithError(c.Writer, 400, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully"})
}
