package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"mime/multipart"
)

type Content struct {
	Title string `json:"title" binding:"required"`
	Story string `json:"story" binding:"required"`
	Date  string `json:"date" binding:"required"`
	Likes int    `json:"likes"`
}

type Response struct {
	Id    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title"`
	Story string             `bson:"story"`
	Date  string             `bson:"date"`
	Likes int                `bson:"likes"`
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

func readCsvFileAndUpdate(form *multipart.FileHeader) error {
	f, err := form.Open()
	if err != nil {
		fmt.Println("Unable to read input file ", err)
		return err
	}
	defer f.Close()
	var r Content
	csvReader := csv.NewReader(f)

	//the following code can be used to replicate transaction so that if theres a error in the file then the transaction gets aborted
	//err = db.Client().UseSession(context.TODO(), func(sessionContext mongo.SessionContext) error {
	//	err := sessionContext.StartTransaction()
	//	var record []string
	//	if err != nil {
	//		sessionContext.AbortTransaction(sessionContext)
	//		return err
	//	}
	//	for {
	//		record, err = csvReader.Read()
	//
	//		if err == io.EOF {
	//			break
	//		}
	//
	//		if err != nil {
	//			break
	//		}
	//
	//		r.Title = record[0]
	//		r.Story = record[1]
	//		r.Date = record[2]
	//		r.Likes = 0
	//		_, err = collection.InsertOne(sessionContext, r)
	//		if err != nil {
	//			break
	//		}
	//	}
	//	if err != nil {
	//		fmt.Println("test")
	//		fmt.Println(err)
	//		sessionContext.AbortTransaction(sessionContext)
	//		return err
	//	} else {
	//		fmt.Println("test2")
	//		fmt.Println(err)
	//		err = sessionContext.CommitTransaction(sessionContext)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	sessionContext.EndSession(sessionContext)
	//	return nil
	//})
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		r.Title = record[0]
		r.Story = record[1]
		r.Date = record[2]
		r.Likes = 0
		_, err = collection.InsertOne(context.TODO(), r)
		if err != nil {
			return err
		}
	}
	return nil
}
