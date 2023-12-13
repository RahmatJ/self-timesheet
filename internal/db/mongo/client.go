package mongo

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func getConfig() config {
	// TODO(Rahmat): read env HOST_NAME, PORT, USER_NAME, PASSWORD, DB_NAME
	return config{
		hostName: "localhost",
		port:     27017,
		userName: "self_timesheet",
		password: "secret",
		dbName:   "self_timesheet_db",
		timeout:  1000,
	}
}

func getUri(cf config) string {
	// TODO(Rahmat): read env IS_LOCAL
	isLocal := true

	baseUri := "mongodb"
	host := fmt.Sprintf("%s:%d", cf.hostName, cf.port)

	if !isLocal {
		baseUri = "mongodb+srv"
		host = cf.hostName
	}

	return fmt.Sprintf("%s://%s:%s@%s/%s", baseUri, cf.userName, cf.password, host, cf.dbName)
}

func InitClient() (*mongo.Client, error) {
	cf := getConfig()
	uri := getUri(cf)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(cf.timeout)*time.Millisecond)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal().Msgf("Error occurred when connect to mongodb. %v", client)
		return nil, err
	}

	defer func() {
		log.Info().Msg("Deferring mongo client")
	}()

	return client, nil
}
