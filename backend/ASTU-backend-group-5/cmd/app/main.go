package main

import (
	"blogApp/internal/config"
	"blogApp/internal/http/routes"
	"blogApp/pkg/mongo"

	"github.com/gin-gonic/gin"
)

func main() {
	confs, err := config.Load()
	if err != nil {
		panic(err)
	}
	mongo.ConnectDB(confs.MONGO_URI)
	defer mongo.DisconnectDB()
	mongo.InitializeCollections()

	userCollection := mongo.UserCollection
	blogCollection := mongo.BlogsCollectionh
	commentCollection := mongo.CommentsCollection
	likeCollection := mongo.LikesCollection
	viewCollection := mongo.ViewsCollection
	tagCollection := mongo.TagsCollection

	router := gin.Default()
	// Serve static files from the "upload" directory
	router.Static("/uploads", "./uploads")
	routes.SetUpRoute(router, blogCollection, commentCollection, likeCollection, viewCollection, tagCollection, userCollection)
	router.Run(":" + confs.GO_PORT)
}
