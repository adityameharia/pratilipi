package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Signup godoc
// @Summary     used to signup a new user
// @Description   the email id password, verifies if the email already exist in the database, if it doesnt exist then it creates a new user and returns an automatically generated userId
// @Tags         signup
// @Accept       json
// @Produce      json
// @Success      200  {object} RespSuccessSignUp
// @Param user body userdata true "get data"
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
	var successResp RespSuccessSignUp
	err = collection.FindOne(context.TODO(), filter).Decode(&resp)
	if err != nil {
		result, err := collection.InsertOne(context.TODO(), body)
		// check for errors in the insertion
		if err != nil {
			c.Writer.WriteHeader(502)
			return
		}
		successResp.Id = result.InsertedID.(primitive.ObjectID).Hex()
		fmt.Println(successResp.Id)
		c.JSON(200, successResp)
		return
	}
	if err == nil {
		if comparePasswords(resp.Password, []byte(orgPassword)) {
			successResp.Id = resp.Id.Hex()
			c.JSON(200, successResp)
		} else {
			c.Writer.WriteHeader(401)
		}
	}
	return
}

// Like godoc
// @Summary     Used to update like
// @Description  Takes the add/remove command from url and updates the like for the respective user
// @Tags         like
// @Produce      json
// @Success      200  {object} RespSuccess
// @Failure      401 {object} RespError
// @Failure      502 {object} RespError
// @Param userid path string true "userid"
// @Param bookid path string true "bookid"
// @Param cmd path string true "add or remove command"
// @Router       /like/{cmd}/{userid}/{bookid} [get]
func Like(c *gin.Context) {
	var errorResp RespError
	var successResp RespSuccess
	_, err := FindUser(c.Param("userId"))
	errorResp.Message = "Invalid UserId"
	if err == mongo.ErrNoDocuments {
		c.JSON(401, errorResp)
		return
	}
	errorResp.Message = "Internal Server Error"
	if err != nil {
		c.JSON(401, errorResp)
		return
	}
	errorResp.Message = "Invalid UserId"
	obId, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		c.JSON(401, errorResp)
		return
	}
	filter := bson.D{primitive.E{Key: "_id", Value: obId}}
	var update bson.M
	if c.Param("cmd") == "remove" {
		update = bson.M{"$pull": bson.M{"liked": c.Param("bookId")}}
	} else {
		update = bson.M{"$push": bson.M{"liked": c.Param("bookId")}}
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		errorResp.Message = "Unable to update your likes"
		c.JSON(502, errorResp)
		return
	}
	successResp.Message = "Liked Updated"
	c.JSON(200, successResp)
	return
}

// FindUserRoute godoc
// @Summary     used to find a user with the given user id
// @Description   used to find a user with the given user id.It returns the entire user document without the hashed password
// @Tags         findUser
// @Produce      json
// @Success      200  {object} RespSuccessFind
// @Failure      401 {object} RespError
// @Failure      502 {object} RespError
// @Param userid path string true "userid"
// @Router       /find/{userid} [get]
func FindUserRoute(c *gin.Context) {
	var errorResp RespError
	var successResp RespSuccessFind
	resp, err := FindUser(c.Param("userId"))
	errorResp.Message = "Invalid UserId"
	if err == mongo.ErrNoDocuments {
		c.JSON(401, errorResp)
		return
	}
	if err != nil {
		errorResp.Message = "Internal Server Error"
		c.JSON(502, errorResp)
		return
	}
	successResp.Message = resp
	c.JSON(200, successResp)
	return
}
