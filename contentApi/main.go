package main

import (
	"context"
	"fmt"
	docs "github.com/adityameharia/pratilipi/contentApi/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	//err := godotenv.Load()
	//if err != nil {
	//	panic(err)
	//}
	docs.SwaggerInfo_swagger.Title = "Content API"
	docs.SwaggerInfo_swagger.Description = "This server responds to the contentApi requests"
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOps := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		panic(err)
	}
	db = client.Database(os.Getenv("DATABASE"))
	defer db.Client().Disconnect(ctx)
	collection = db.Collection(os.Getenv("COLLECTION"))
	count, err = collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	//docs.SwaggerInfo_swagger.BasePath = "/"
	r.Use(CORSMiddleware())
	r.POST("/csv/:userid", ValidateUser(), ReadCSV)
	r.POST("/like/:cmd/:userid/:bookid", ValidateUser(), Like)
	r.GET("/getmostliked/:userid", ValidateUser(), GetMostLiked)
	r.GET("/books/:userid/:pageno", ValidateUser(), GetBooks)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + os.Getenv("PORT"))
}
