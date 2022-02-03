package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/mail"
	"time"
	"unicode"
)

type userdata struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type response struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 10 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

// @BasePath /
// Ping-example godoc
// @Summary illustrates ping examples
// @Description returns pong on ping
// @Accept json
// @Produce json
// @Router /ping [get]
func signup(c *gin.Context) {
	var body userdata
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !isValidEmail(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}
	if isValidPassword(body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	hashedPwd := string(hash)
	user := bson.D{{"email", body.Email}, {"password", hashedPwd}}

	filter := bson.D{primitive.E{Key: "email", Value: body.Email}}
	var resp response
	err = collection.FindOne(context.TODO(), filter).Decode(&resp)
	if err != nil {
		result, err := collection.InsertOne(context.TODO(), user)
		// check for errors in the insertion
		if err != nil {
			c.JSON(502, gin.H{
				"message": "Unable to add email and password to database",
			})
			return
		}
		c.JSON(200, gin.H{
			"id": result.InsertedID,
		})
		return
	}
	if err == nil {
		if comparePasswords(resp.Password, []byte(body.Password)) {
			fmt.Println(resp)
			c.JSON(200, gin.H{
				"id": resp.Id,
			})
		} else {
			c.JSON(401, gin.H{
				"message": "The email Id you have entered has already been registered and the password you are entering is wrong",
			})
		}
	}
	return
}

func validateUser(c *gin.Context) {
	time.Sleep(10 * time.Second)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
