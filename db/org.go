package db

import (
	"context"
	"talentprotocol/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) InsertOrganization(org *types.Organization) error {
	_, err := db.organizationsCollection.InsertOne(context.TODO(), org)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetOrgDetails(orgLogin *types.OrgLogin) (*types.Organization, error) {
	response := &types.Organization{}
	err := db.organizationsCollection.FindOne(context.TODO(), bson.M{"org_email": orgLogin.Email}).Decode(response)
	if err != nil {
		return nil, err
	}

	// if !db.ValidatePassword(orgLogin.Password, orgLogin.Password) {
	// 	return nil, errors.New("org not authenticated")
	// }

	return response, err
}

func (db *DB) CreateJobOpening(org *types.JobOpening) (any, error) {
	res, err := db.orgOpeningsCollection.InsertOne(context.TODO(), org)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (db *DB) GetAllJobOpenings(orgName string) ([]*types.JobOpening, error) {
	pipeline := bson.M{
		"org_name": orgName,
	}

	cursor, err := db.orgOpeningsCollection.Find(context.TODO(), pipeline)

	if err != nil {
		return nil, err
	}

	allOpenings := []*types.JobOpening{}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var ope *types.JobOpening
		if err := cursor.Decode(&ope); err != nil {
			return nil, err
		}
		allOpenings = append(allOpenings, ope)
	}

	return allOpenings, nil
}
