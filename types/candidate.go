package types

import "time"

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
	CandidateEmail     string    `json:"candidate_email" bson:"candidate_email"`
	JobOpeningID       string    `json:"job_opening_id" bson:"job_opening_id"`
	JobApplicationDate time.Time `json:"job_application_date" bson:"job_application_date"`
	Status             string    `json:"status" bson:"status"`
}
