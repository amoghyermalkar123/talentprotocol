package db

import (
	"context"
	"errors"
	"talentprotocol/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) AddCandidateDetails(userDetails *types.CandidateDetails) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userDetails.Password = string(hashedPwd)

	_, err = db.candidateCollection.InsertOne(context.TODO(), userDetails)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetCandidateDetails(userLogin *types.CandidateLogin) (*types.CandidateDetails, error) {
	userResponse := &types.CandidateDetails{}
	err := db.candidateCollection.FindOne(context.TODO(), bson.M{"email": userLogin.Email}).Decode(userResponse)
	if err != nil {
		return nil, err
	}

	if !db.ValidatePassword(userLogin.Password, userResponse.Password) {
		return nil, errors.New("user not authenticated")
	}

	response := &types.CandidateDetails{
		FullName: userResponse.FullName,
		Email:    userResponse.Email,
		Age:      userResponse.Age,
	}

	return response, err
}

func (db *DB) ValidatePassword(inputPassword, dbPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword)); err != nil {
		return false
	}
	return true
}

func (db *DB) GetJobOpeningsNotAppliedTo(candidateEmail string) ([]*types.JobOpening, error) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	// Find job openings posted in the last 7 days
	filter := bson.M{
		"job_posted_at": bson.M{"$gte": sevenDaysAgo},
	}

	// Get a list of job opening IDs that the candidate has applied to
	var appliedOpeningIDs []primitive.ObjectID
	cursor, err := db.candidateJobApplicationsCollection.Find(context.TODO(), bson.M{"candidate_email": candidateEmail})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var app *types.CandidateJobApplication
		if err := cursor.Decode(&app); err != nil {
			return nil, err
		}
		appliedOpeningIDs = append(appliedOpeningIDs, app.JobOpeningID)
	}

	if len(appliedOpeningIDs) > 0 {
		// Find job openings not in the list of appliedOpeningIDs
		filter["_id"] = bson.M{"$nin": appliedOpeningIDs}
	}

	// Execute the query
	cursor, err = db.orgOpeningsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var unappliedOpenings []*types.JobOpening
	for cursor.Next(context.TODO()) {
		var opening *types.JobOpening
		if err := cursor.Decode(&opening); err != nil {
			return nil, err
		}
		unappliedOpenings = append(unappliedOpenings, opening)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return unappliedOpenings, nil
}

func (db *DB) InsertCandidateSubmission(candEmail, jobOpeningID string, submission *types.CandidateSubmission) error {
	jobID, err := primitive.ObjectIDFromHex(jobOpeningID)
	if err != nil {
		return err
	}

	db.insertCandidateJobApplication(&types.CandidateJobApplication{
		CandidateEmail:     candEmail,
		JobOpeningID:       jobID,
		JobApplicationDate: time.Now(),
		Status:             Open.String(),
		Answers:            submission.Answers,
	})
	return nil
}

func (db *DB) insertCandidateJobApplication(application *types.CandidateJobApplication) error {
	_, err := db.candidateJobApplicationsCollection.InsertOne(context.TODO(), application)
	if err != nil {
		return err
	}
	return nil
}
