package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"mime/multipart"
)

type Content struct {
	Title string `json:"title" binding:"required"`
	Story string `json:"story" binding:"required"`
	Date  string `json:"date" binding:"required"`
	Likes int64  `json:"likes"`
}

type Response struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
	Story string             `json:"story" bson:"story"`
	Date  string             `json:"date" bson:"date"`
	Likes int                `json:"likes" bson:"likes"`
	Liked bool               `json:"liked" bson:"liked"`
	Count int64              `json:"count" bson:"count"`
}

type ResponseWithCount struct {
	Response []Response `json:"data"`
	Count    int64      `json:"count"`
}

type Userdata struct {
	Id       string   `json:"Id" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Liked    []string `json:"liked" binding:"required"`
}
type Mess struct {
	Message Userdata `json:"message" binding:"required"`
}

func readCsvFileAndUpdate(form *multipart.FileHeader) (int64, error) {
	f, err := form.Open()
	if err != nil {
		fmt.Println("Unable to read input file ", err)
		return count, err
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
			return count, err
		}

		r.Title = record[0]
		r.Story = record[1]
		r.Date = record[2]
		r.Likes = 0
		_, err = collection.InsertOne(context.TODO(), r)
		if err != nil {
			return count, err
		}
	}
	count, err := collection.CountDocuments(context.TODO(), bson.D{})
	return count, nil
}

func fillLiked(resp *[]Response, field Mess) {
	fmt.Println("test call")
	fmt.Println(field.Message.Id)
	fmt.Println(field.Message.Liked)
	for i, r := range *resp {
		(*resp)[i].Liked = Find(field.Message.Liked, primitive.ObjectID.Hex(r.Id))
	}

}

func Find(slice []string, val string) bool {
	//fmt.Println(val)
	for _, item := range slice {
		if item == val {
			fmt.Println("testinng if matched")
			return true
		}
	}
	return false
}
