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
)

// @BasePath /
// Ping-example godoc
// @Summary illustrates ping examples
// @Description returns pong on ping
// @Accept json
// @Produce json
// @Router /ping [get]
func signup(c *gin.Context) {
	var orgPassword string
	var body userdata
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgPassword = body.Password
	body.Liked = make([]string, 0)
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
	body.Password = hashedPwd
	//user := bson.D{{"email", body.Email}, {"password", hashedPwd}, {"likedBooks", bson.A{}}}

	filter := bson.D{primitive.E{Key: "email", Value: body.Email}}
	var resp response
	err = collection.FindOne(context.TODO(), filter).Decode(&resp)
	if err != nil {
		result, err := collection.InsertOne(context.TODO(), body)
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
		if comparePasswords(resp.Password, []byte(orgPassword)) {
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

func addLike(c *gin.Context) {
	if !validateUser(c.Param("userId")) {
		c.JSON(401, gin.H{
			"message": "Invalid UserId",
		})
		return
	}
	obId, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid UserId",
		})
		return
	}
	filter := bson.D{primitive.E{Key: "_id", Value: obId}}

	update := bson.M{"$push": bson.M{"liked": c.Param("bookId")}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(502, gin.H{
			"message": "Unable to update your likes",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Liked Updated",
	})
	return
}
