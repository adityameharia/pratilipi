package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
	"unicode"
)

type userdata struct {
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Liked    []string `json:"liked" example:[]`
}

type signUpBodySwagger struct {
	Email    string `json:"email" binding:"required" example:"adityameh@gmail.com"`
	Password string `json:"password" binding:"required" example:"Qwert@009"`
}

type response struct {
	Id       primitive.ObjectID `bson:"_id, omitempty" example:"507c7f79bcf86cd7994f6c0e"`
	Email    string             `json:"email" binding:"required" example:"adi@gmail.com"`
	Password string             `json:"password" binding:"required" example:""`
	Liked    []string           `json:"liked" binding:"required"`
}

type RespError struct {
	Message string `json:"message" binding:"required" example:"Error"`
}

type RespSuccess struct {
	Message string `json:"message" binding:"required" example:"Data Updated Successfully"`
}

type responseFind struct {
	Id    primitive.ObjectID `bson:"_id, omitempty" example:"507c7f79bcf86cd7994f6c0e"`
	Email string             `json:"email" binding:"required" example:"adi@gmail.com"`
	Liked []string           `json:"liked" binding:"required" example:"[507f1f77bcf86cd799439011,507f1f77bcf86cd799439011]"`
}

type RespSuccessFind struct {
	Message responseFind `json:"message" binding:"required"`
}

type RespSuccessSignUp struct {
	Id string `json:"id" binding:"required" example:"507c7f79bcf86cd7994f6c0e"`
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

func FindUser(id string) (responseFind, error) {
	var resp responseFind
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return responseFind{}, err
	}
	opts := options.FindOne().SetProjection(bson.D{{"password", 0}})
	err = collection.FindOne(context.TODO(), bson.D{{"_id", obId}}, opts).Decode(&resp)
	if err != nil {
		fmt.Println(err)
		return responseFind{}, err
	}
	return resp, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
