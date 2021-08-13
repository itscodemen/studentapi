package handlers

import (
	"log"
	"proj1/api/auth"
	"proj1/api/utils"
	"proj1/domain/storage"
	"proj1/models"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, gin.H{"Message": "invalid json"})
		return
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(500, gin.H{"Message": "error hashing password"})
		return
	}

	err = storage.User.CreateUser(user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"Message": "error creating user"})
		return
	}

	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var payload models.LoginPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Message": "invalid json"})
		return
	}

	if !CheckUserExists(c, payload.Email) {
		log.Println(payload.Email)
		utils.RespondWithError(c.Writer, 401, "Invalid User credentials")
		return
	}

	if CheckPassword(c, payload.Password) {
		utils.RespondWithError(c.Writer, 401, "Invalid User credentials")
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"Message": "Error signing token"})
		return
	}

	tokenResponse := models.LoginResponse{
		Token: signedToken,
	}

	c.JSON(200, tokenResponse)

}

func CheckUserExists(c *gin.Context, email string) bool {
	_, err := storage.User.CheckUserExists(email)
	return err == nil
}

func CheckPassword(c *gin.Context, password string) bool {
	err := storage.User.CheckPassword(password)
	return err == nil
}
