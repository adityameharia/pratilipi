package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	err = readCsvFileAndUpdate(file)
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
	opts := options.Find().SetSort(sort).SetLimit(10)
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
	c.JSON(200, gin.H{
		"mostLiked": resp,
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
	opts := options.Find().SetSkip(pgno * 10).SetLimit(10)
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
	c.JSON(200, gin.H{
		"books": resp,
	})
	return
}
