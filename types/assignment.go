package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrgAssignments struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	AssignmentName        string             `bson:"assignment_name"`
	CodeProblemStatement  string             `bson:"code_problem_statement"`
	TechnicalQuestions    []string           `bson:"technical_questions"`
	CodeAnalysisQuestions []string           `bson:"code_analysis_questions"`
}

type CandidateSubmission struct {
	CandidateID         string             `bson:"candidate_id"`
	SubmittedAssignment primitive.ObjectID `bson:"submitted_assignment"`
	Answers             struct {
		Code                string `bson:"code"`
		Rating              string `bson:"rating"`
		CodeAnalysisAnswers string `bson:"code_analysis_answers"`
		TechnicalAnswers    string `bson:"technical_answers"`
	} `bson:"answers"`
}
