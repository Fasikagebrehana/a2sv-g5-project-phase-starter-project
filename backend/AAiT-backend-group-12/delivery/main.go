package main

import (
	"blog_api/delivery/env"
	"blog_api/delivery/router"
	initdb "blog_api/infrastructure/db"
	google_auth "blog_api/infrastructure/oauth"
	redis_service "blog_api/infrastructure/redis"
	"fmt"
	"log"
)

func main() {
	// load environment variables
	err := env.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// connect to mongodb
	mongoClient, err := initdb.ConnectDB(env.ENV.DB_ADDRESS, env.ENV.DB_NAME)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	database := mongoClient.Database(env.ENV.DB_NAME)

	err = initdb.CreateRootUser(database, env.ENV.ROOT_USERNAME, env.ENV.ROOT_PASSWORD)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer initdb.DisconnectDB(mongoClient)

	// connect to regis
	redisClient, err := redis_service.ConnectStore(env.ENV.REDIS_URL)
	if err != nil {
		log.Println(err)
	}

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatalln("Redis connection was refused: " + err.Error())
		return
	}

	defer redisClient.Close()

	// create google provider for oauth
	google_auth.NewAuth(env.ENV.GOOGLE_CLIENT_ID, env.ENV.GOOGLE_CLIENT_SECRET, 1, fmt.Sprintf("http://localhost:%v/auth/google/callback", env.ENV.PORT))

	// setup router
	router.SetupRouter(env.ENV.PORT, env.ENV.ROUTE_PREFIX, database, redisClient)
}
