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

func readCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	if extension != ".csv" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "File Extension not available",
		})
		return
	}
	count, err = readCsvFileAndUpdate(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data added successfully",
	})
}

func getMostLiked(c *gin.Context) {
	sort := bson.D{{"likes", -1}}
	opts := options.Find().SetSort(sort).SetLimit(9)
	cur, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		fmt.Println(err)
		c.JSON(502, gin.H{
			"message": "Unable to access database",
		})
		return
	}
	var res []Response
	if err = cur.All(context.TODO(), &res); err != nil {
		fmt.Println(err)
		c.JSON(502, gin.H{
			"message": "Unable to access database",
		})
		return
	}
	user, _ := c.Get("user")
	field, _ := user.(UserApiResponse)
	fillLiked(&res, field)
	c.JSON(200, gin.H{
		"mostLiked": res,
	})
	return
}

func getBooks(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("pageno"))
	pgno := int64(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Page Number",
		})
	}
	opts := options.Find().SetSkip((pgno - 1) * 10).SetLimit(9)
	cur, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		fmt.Println(err)
		c.JSON(502, gin.H{
			"message": "Unable to access database",
		})
		return
	}
	var resp []Response
	if err = cur.All(context.TODO(), &resp); err != nil {
		fmt.Println(err)
		c.JSON(502, gin.H{
			"message": "Unable to access database",
		})
		return
	}

	user, _ := c.Get("user")
	field, _ := user.(UserApiResponse)
	fillLiked(&resp, field)
	var respWCount ResponseWithCount
	respWCount.Count = count
	respWCount.Response = resp
	c.JSON(200, gin.H{
		"books": respWCount,
	})
	return
}

func Like(c *gin.Context) {
	user, _ := c.Get("user")
	userData, _ := user.(UserApiResponse)
	obId, err := primitive.ObjectIDFromHex(c.Param("bookid"))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid UserId",
		})
		return
	}
	command := Find(userData.Message.Liked, c.Param("bookid"))
	var update bson.M
	var url string
	if !command {
		update = bson.M{"$inc": bson.M{"likes": 1}}
	} else {
		update = bson.M{"$inc": bson.M{"likes": -1}}
	}
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"_id", obId}}, update)
	if err != nil {
		c.JSON(502, gin.H{
			"message": "Unable to update your likes",
		})
		return
	}
	if !command {
		url = "http://localhost:8000/like/add/" + c.Param("userid") + "/" + c.Param("bookid")
	} else {
		url = "http://localhost:8000/like/remove/" + c.Param("userid") + "/" + c.Param("bookid")
	}
	_, err = http.Get(url)
	if err != nil {
		fmt.Println(err)
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
