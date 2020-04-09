package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	env "github.com/joho/godotenv"

	"product-service/api/rest"
	"product-service/core/repo"
	"product-service/middleware/business"
	"product-service/middleware/storage/mongo"
	"product-service/middleware/storage/postgres"
	"product-service/middleware/storage/redis"
	"product-service/server"
)

func main() {
	// loadEnv()
	repo := getRepository()
	if repo == nil {
		return
	}
	service := business.NewSneakerProductService(repo)
	handler := rest.NewHandler(service, os.Getenv("CONTENT_TYPE"))
	routes := rest.ProvideRoutes(handler)
	srv := server.NewInstance(os.Getenv("HOST"))
	srv.SetupRouter(routes)
	srv.Start()
}

func loadEnv() {
	err := env.Load("./env/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func httpPort() string {
	port := "8420"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func getRepository() repo.SneakerProductRepository {
	switch os.Getenv("USE_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := redis.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		mongoCollection := os.Getenv("MONGO_COLLECTION")
		repo, err := mongo.NewMongoRepository(
			mongoURL,
			mongodb,
			mongoCollection,
			mongoTimeout,
		)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		postgresUrl := os.Getenv("POSTGRES_URL")
		postgresTable := os.Getenv("POSTGRES_TABLE")
		repo, err := postgres.NewPostgresRepository(postgresUrl, postgresTable)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
