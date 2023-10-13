package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Organization struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	OrgName        string               `bson:"org_name" json:"org_name"`
	Assignments    []primitive.ObjectID `bson:"assignments" json:"assignments"`
	SubscriptionID string               `bson:"subscription_id" json:"subscription_id"`
	OrgEmail       string               `bson:"org_email" json:"org_email"`
}

type JobOpening struct {
	ID             string    `bson:"_id,omitempty" json:"_id"`
	OrganizationID string    `bson:"org_id" json:"org_id"`
	OrgName        string    `bson:"org_name" json:"org_name"`
	OpeningName    string    `bson:"opening_name" json:"opening_name"`
	JobDescription string    `bson:"jd" json:"jd"`
	JobPostedAt    time.Time `bson:"job_posted_at" json:"job_posted_at,omitempty"`
}
