package db

import (
	"context"
	"talentprotocol/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (db *DB) CreateJobOpening(org *types.JobOpening) error {
	_, err := db.orgOpeningsCollection.InsertOne(context.TODO(), org)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetAllJobOpenings(orgName string) ([]*types.OrgOpeningsWithAssigs, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"org_name": orgName},
		},
		{
			"$lookup": bson.M{
				"from":         "org_assignments",
				"localField":   "_id",
				"foreignField": "opening_id",
				"as":           "assignments",
			},
		},
	}

	cursor, err := db.orgOpeningsCollection.Aggregate(context.TODO(), pipeline)

	if err != nil {
		return nil, err
	}

	allOpenings := []*types.OrgOpeningsWithAssigs{}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var ope *types.OrgOpeningsWithAssigs
		if err := cursor.Decode(&ope); err != nil {
			return nil, err
		}
		allOpenings = append(allOpenings, ope)
	}

	return allOpenings, nil
}

func (db *DB) GetAssignmentForOrgOpening(openingID string) (*types.OrgAssignments, error) {
	objectID, err := primitive.ObjectIDFromHex(openingID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"opening_id": objectID}
	var assignment *types.OrgAssignments

	err = db.orgAssignmentsCollection.FindOne(context.TODO(), filter).Decode(&assignment)
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (db *DB) InsertOrgAssignment(assignment *types.OrgAssignments, openingID string) error {
	objectID, err := primitive.ObjectIDFromHex(openingID)
	if err != nil {
		return err
	}

	assignment.JobOpeningID = objectID
	_, err = db.orgAssignmentsCollection.InsertOne(context.TODO(), assignment)
	if err != nil {
		return err
	}

	return nil
}
