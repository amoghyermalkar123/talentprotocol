package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrgAssignments struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	JobOpeningID          primitive.ObjectID `bson:"opening_id" json:"opening_id,omitempty"`
	AssignmentName        string             `bson:"assignment_name" json:"assignment_name"`
	CodeProblemStatement  string             `bson:"code_problem_statement" json:"code_problem_statement"`
	TechnicalQuestions    []string           `bson:"technical_questions" json:"technical_questions"`
	CodeAnalysisQuestions []string           `bson:"code_analysis_questions" json:"code_analysis_questions"`
}

type CandidateSubmission struct {
	SubmissionBy          string `bson:"submission_by"`
	SubmittedAssignmentID string `bson:"submitted_assignment_id" json:"submitted_assignment_id"`
	Answers               struct {
		Code                string   `bson:"code" json:"code"`
		Rating              int      `bson:"rating" json:"rating,omitempty"`
		CodeAnalysisAnswers []string `bson:"code_analysis_answers" json:"code_analysis_answers,omitempty"`
		TechnicalAnswers    []string `bson:"technical_answers" json:"technical_answers,omitempty"`
	} `bson:"answers" json:"answers"`
}
