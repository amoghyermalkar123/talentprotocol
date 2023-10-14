package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client                             *mongo.Client
	Log                                zerolog.Logger
	candidateCollection                *mongo.Collection
	organizationsCollection            *mongo.Collection
	orgOpeningsCollection              *mongo.Collection
	orgAssignmentsCollection           *mongo.Collection
	candidateJobApplicationsCollection *mongo.Collection
	candidateSubmissionsCollection     *mongo.Collection
}

func GetDB(host string) (*DB, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017/talentprotocol", host)))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &DB{
		client:                             client,
		candidateCollection:                client.Database("talentprotocol").Collection("candidate"),
		organizationsCollection:            client.Database("talentprotocol").Collection("organizations"),
		orgOpeningsCollection:              client.Database("talentprotocol").Collection("org_job_openings"),
		orgAssignmentsCollection:           client.Database("talentprotocol").Collection("org_assignments"),
		candidateJobApplicationsCollection: client.Database("talentprotocol").Collection("candidate_job_applications"),
		candidateSubmissionsCollection:     client.Database("talentprotocol").Collection("candidate_submissions"),
	}, nil
}
