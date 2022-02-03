package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
	"unicode"
)

type userdata struct {
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Liked    []string `json:"liked"`
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

func validateUser(id string) bool {
	fmt.Println(id)
	var resp response
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = collection.FindOne(context.TODO(), bson.D{{"_id", obId}}).Decode(&resp)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
