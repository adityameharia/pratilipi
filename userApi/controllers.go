package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Signup godoc
// @Summary     used to signup a new user
// @Description  takes the email id password, verifies if the email already exist in the database, if it doesnt exist then it creates a new user and returns an automatically generated userId
// @Tags         signup
// @Accept       json
// @Produce      json
// @Success      200  {object} resp
// @Failure      401  {object} resp
// @Failure      502  {object} resp
// @Router       /signup [post]
func Signup(c *gin.Context) {
	var orgPassword string
	var body userdata
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(400)
		return
	}
	orgPassword = body.Password
	body.Liked = make([]string, 0)
	body.Liked = append(body.Liked, "61ff66fa02cd037ea1ea0bc3")
	if !isValidEmail(body.Email) {
		c.Writer.WriteHeader(400)
		return
	}
	if !isValidPassword(body.Password) {
		c.Writer.WriteHeader(400)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	hashedPwd := string(hash)
	body.Password = hashedPwd
	filter := bson.D{primitive.E{Key: "email", Value: body.Email}}
	var resp response
	err = collection.FindOne(context.TODO(), filter).Decode(&resp)
	if err != nil {
		result, err := collection.InsertOne(context.TODO(), body)
		// check for errors in the insertion
		if err != nil {
			c.Writer.WriteHeader(502)
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
			c.Writer.WriteHeader(401)
		}
	}
	return
}

// AddLike godoc
// @Summary     likes a book for the particular user
// @Description  takes the userId and bookId from the parameters. Validates if the user exists and then adds the book to the list of books the user has liked
// @Tags         AddLike
// @Accept       json
// @Produce      json
// @Param        userId  path  string  true  "User ID"
// @Param        bookId  path  string  true  "Book ID"
// @Success      200  {object} resp
// @Failure      401  {object} resp
// @Failure      502  {object} resp
// @Router       /like/:userId/:bookId [post]
func AddLike(c *gin.Context) {
	_, err := FindUser(c.Param("userId"))
	if err == mongo.ErrNoDocuments {
		c.JSON(401, gin.H{
			"message": "Invalid UserId",
		})
		return
	}
	if err != nil {
		c.JSON(502, gin.H{
			"message": "Internal Server Error",
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

func RemoveLike(c *gin.Context) {
	_, err := FindUser(c.Param("userId"))
	if err == mongo.ErrNoDocuments {
		c.JSON(401, gin.H{
			"message": "Invalid UserId",
		})
		return
	}
	if err != nil {
		c.JSON(502, gin.H{
			"message": "Internal Server Error",
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

	update := bson.M{"$pull": bson.M{"liked": c.Param("bookId")}}
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

func FindUserRoute(c *gin.Context) {
	resp, err := FindUser(c.Param("userId"))
	if err == mongo.ErrNoDocuments {
		c.JSON(401, gin.H{
			"message": "Invalid UserId",
		})
		return
	}
	if err != nil {
		c.JSON(502, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": resp,
	})
	return
}
