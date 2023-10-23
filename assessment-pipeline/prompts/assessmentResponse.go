package prompts

import (
	"encoding/json"
	"fmt"
)

type EvaluationOutput struct {
	Rating            string `json:"rating"`
	CodeAnalysis      string `json:"code_analysis"`
	EvaluationSummary string `json:"evaluation_summary"`
}

type EvaluationInput struct {
	RequirementFactors []string
	Code               string
	CodeQnA            map[string]string
	TechnicalQnA       map[string]string
}

func GetOutputFormat() string {
	eo := &EvaluationOutput{}
	str, _ := json.Marshal(eo)
	return string(str)
}

func GenerateEvaluationPrompt(evalInput *EvaluationInput) string {
	var EvaluationPrompt string = fmt.Sprintf(
		`analyse the following code %s, the following technical QnA %v and
		these QNA based on the code %v
		Evaluate the candidate over the following requirements %v and finally provide
		your evaluation in the following JSON format %s.
	`,
		evalInput.Code,
		evalInput.TechnicalQnA,
		evalInput.CodeQnA,
		evalInput.RequirementFactors,
		GetOutputFormat(),
	)

	return EvaluationPrompt
}
