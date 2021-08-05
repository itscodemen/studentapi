package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
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

func Validation(c *gin.Context, name string, email string, phone string) (bool, string) {
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
