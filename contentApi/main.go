package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var collection *mongo.Collection
var db *mongo.Database
var count int64

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOps := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("successfully connected")
	}
	db = client.Database(os.Getenv("DATABASE"))
	defer db.Client().Disconnect(ctx)
	collection = db.Collection(os.Getenv("COLLECTION"))
	count, err = collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(ValidateUser())
	r.POST("/csv/:userid", readCSV)
	r.POST("/like/:userid/:bookid", Like)
	r.GET("/getmostliked/:userid", getMostLiked)
	r.GET("/books/:userid/:pageno", getBooks)
	r.Run(os.Getenv("PORT"))
}
