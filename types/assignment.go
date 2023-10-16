package types

type CandidateSubmission struct {
	Answers struct {
		Code                string   `bson:"code" json:"code"`
		Rating              int      `bson:"rating" json:"rating,omitempty"`
		CodeAnalysisAnswers []string `bson:"code_analysis_answers" json:"code_analysis_answers,omitempty"`
		TechnicalAnswers    []string `bson:"technical_answers" json:"technical_answers,omitempty"`
	} `bson:"answers" json:"answers"`
}
