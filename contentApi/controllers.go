package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"path/filepath"
	"strconv"
)

// ReadCSV godoc
// @Summary     parse csv file and update data to the database
// @Description  parse csv file and update data to the database
// @Tags         csv
// @Accept 		multipart/form-data
// @Produce      json
// @Success      200  {object} RespSuccess
// @Failure      400 {object} RespError
// @Failure      502 {object} RespError
// @Param userid path string true "userid"
// @Param file formData file true "file"
// @Router       /csv/{userid} [post]
func ReadCSV(c *gin.Context) {
	var errorResp RespError
	var successResp RespSuccess
	errorResp.Message = "No file is received"
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp)
		return
	}
	extension := filepath.Ext(file.Filename)
	errorResp.Message = "File Extension not available"
	if extension != ".csv" {
		c.JSON(http.StatusBadRequest, errorResp)
		return
	}
	count, err = readCsvFileAndUpdate(file)
	errorResp.Message = err.Error()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp)
		return
	}
	successResp.Message = "Data added successfully"
	c.JSON(http.StatusOK, successResp)
}

// GetMostLiked godoc
// @Summary     finds and returns top content
// @Description   finds and returns top content on the basis of number of likes
// @Tags         TopContent
// @Produce      json
// @Success      200  {object} RespSuccessML
// @Failure      502 {object} RespError
// @Param userid path string true "userid"
// @Router       /getmostliked/{userid} [get]
func GetMostLiked(c *gin.Context) {
	var errorResp RespError
	var successResp RespSuccessML
	sort := bson.D{{"likes", -1}}
	opts := options.Find().SetSort(sort).SetLimit(9)
	cur, err := collection.Find(context.TODO(), bson.D{}, opts)
	errorResp.Message = "Unable to access database"
	if err != nil {
		fmt.Println(err)
		c.JSON(502, errorResp)
		return
	}
	var res []Response
	if err = cur.All(context.TODO(), &res); err != nil {
		fmt.Println(err)
		c.JSON(502, errorResp)
		return
	}
	user, _ := c.Get("user")
	field, _ := user.(UserApiResponse)
	fillLiked(&res, field)
	successResp.MostLiked = res
	c.JSON(200, successResp)
	return
}

// GetBooks godoc
// @Summary     finds and returns books
// @Description   finds and returns books
// @Tags         Books
// @Produce      json
// @Success      200  {object} RespSuccessBooks
// @Failure      502 {object} RespError
// @Param userid path string true "userid"
// @Param pageno path string true "pageno"
// @Router       /books/{userid}/{pageno} [get]
func GetBooks(c *gin.Context) {
	var errorResp RespError
	var successResp RespSuccessBooks
	i, err := strconv.Atoi(c.Param("pageno"))
	errorResp.Message = "Invalid Page Number"
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp)
	}
	pgno := int64(i)
	opts := options.Find().SetSkip((pgno - 1) * 10).SetLimit(9)
	cur, err := collection.Find(context.TODO(), bson.D{}, opts)
	errorResp.Message = "Unable to access database"
	if err != nil {
		c.JSON(502, errorResp)
		return
	}
	var resp []Response
	if err = cur.All(context.TODO(), &resp); err != nil {
		fmt.Println(err)
		c.JSON(502, errorResp)
		return
	}

	user, _ := c.Get("user")
	field, _ := user.(UserApiResponse)
	fillLiked(&resp, field)
	var respWCount ResponseWithCount
	respWCount.Count = count
	respWCount.Response = resp
	successResp.Books = respWCount
	c.JSON(200, successResp)
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
// @Router       /like/{cmd}/{userid}/{bookid} [post]
func Like(c *gin.Context) {
	var errorResp RespError
	var successResp RespSuccess
	obId, err := primitive.ObjectIDFromHex(c.Param("bookid"))
	errorResp.Message = "Invalid UserId"
	if err != nil {
		c.JSON(401, errorResp)
		return
	}
	command := c.Param("cmd")
	var update bson.M
	var url string
	if command == "add" {
		update = bson.M{"$inc": bson.M{"likes": 1}}
	} else {
		update = bson.M{"$inc": bson.M{"likes": -1}}
	}
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"_id", obId}}, update)
	errorResp.Message = "Unable to update your likes"
	if err != nil {
		c.JSON(502, errorResp)
		return
	}
	url = "http://localhost:8000/like/" + c.Param("cmd") + "/" + c.Param("userid") + "/" + c.Param("bookid")
	_, err = http.Get(url)
	if err != nil {
		fmt.Println(err)
		c.JSON(502, errorResp)
		return
	}
	successResp.Message = "Liked Updated"
	c.JSON(200, successResp)
	return
}
