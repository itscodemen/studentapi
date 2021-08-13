package utils

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RespondWithJSON(w http.ResponseWriter, reqcode int, eh interface{}) error {
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

func RespondWithError(w http.ResponseWriter, code int, msg string) error {
	return RespondWithJSON(w, code, map[string]string{"error": msg})
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func isNameValid(e string) bool {
	nameRegex := regexp.MustCompile(`^(\D)([A-Z]|[a-z]|[ ]|[0-9]|[.]){1,64}`)
	return nameRegex.MatchString(e)
}

func Validation(c *gin.Context, name string, email string, phone string) (bool, string) {
	if !isNameValid(name) {
		return true, "Invalid Name"
	}
	if !isEmailValid(email) {
		return true, "Invalid Email"
	}
	if len(phone) != 10 {
		return true, "Please Enter 10 Digit Phone Number"
	}
	return false, ""
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

// func CheckPassword(providedPassword string) error {
// 	var user models.User
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
