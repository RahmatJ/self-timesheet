package timesheet

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryImpl struct {
	collection *mongo.Collection
}

func (r *repositoryImpl) Create(ctx context.Context, payload CreateTimesheetRecordPayload) error {
	log.Info().Msgf("Inserting timesheet data %v", payload)

	_, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		log.Error().Msgf("Error happened when create insert data %v", err)
		return err
	}
	log.Info().Msg("Succeed insert timesheet")

	return nil
}

func NewRepository(database *mongo.Database) Repository {
	collectionName := "timesheet"
	collection := database.Collection(collectionName)

	return &repositoryImpl{
		collection: collection,
	}
}
