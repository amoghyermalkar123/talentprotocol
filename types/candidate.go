package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CandidateDetails struct {
	FullName string `json:"full_name" bson:"full_name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Age      string `json:"age" bson:"age"`
}

type CandidateLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CandidateJobApplication struct {
	CandidateEmail     string             `json:"candidate_email" bson:"candidate_email"`
	JobOpeningID       primitive.ObjectID `json:"job_opening_id" bson:"job_opening_id"`
	JobApplicationDate time.Time          `json:"job_application_date" bson:"job_application_date"`
	Status             string             `json:"status" bson:"status"`
	Answers            struct {
		Code                string   `bson:"code" json:"code"`
		Rating              int      `bson:"rating" json:"rating,omitempty"`
		CodeAnalysisAnswers []string `bson:"code_analysis_answers" json:"code_analysis_answers,omitempty"`
		TechnicalAnswers    []string `bson:"technical_answers" json:"technical_answers,omitempty"`
	} `bson:"answers" json:"answers"`
}
