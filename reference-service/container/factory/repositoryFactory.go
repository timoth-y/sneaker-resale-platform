package factory

import (
	"log"

	"go.kicksware.com/api/reference-service/core/repo"
	"go.kicksware.com/api/reference-service/env"
	"go.kicksware.com/api/reference-service/usecase/storage/mongo"
	"go.kicksware.com/api/reference-service/usecase/storage/postgres"
)

func ProvideRepository(config env.ServiceConfig) repo.SneakerReferenceRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewMongoRepository(config.Mongo); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewPostgresRepository(config.Postgres); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
