package assessmentpipeline

import (
	"encoding/json"
	"fmt"
	"log"
	"talentprotocol/assessment-pipeline/prompts"
	"talentprotocol/db"
	"talentprotocol/types"

	"github.com/nats-io/nats.go"
)

// AssessmentPipeline describes the assessment process
// receives assignments, evaluates them and
type AssessmentPipeline struct {
	assignments      chan types.CandidateSubmission
	consumptionQueue uint32
	natsHost         string
	Db               *db.DB
	EvaluationTopic  string
	Nats             *nats.Conn
}

func NewAssessmentPipeline(db *db.DB) *AssessmentPipeline {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Printf("Error connecting to NATS: %v", err)
	}

	return &AssessmentPipeline{
		assignments:      make(chan types.CandidateSubmission),
		consumptionQueue: 64,
		Db:               db,
		EvaluationTopic:  "eval-queue",
		Nats:             nc,
	}
}

// SetupAssessmentPipeline
func (ap *AssessmentPipeline) AssessmentPipeline() {
	ap.assignments = make(chan types.CandidateSubmission)

	go ap.assignmentConsumer()
	for {
		select {
		case evaluation := <-ap.assignments:
			technicalQnA := make(map[string]string)
			codeQnA := make(map[string]string)
			for idx, question := range evaluation.Assignment.TechnicalQuestions {
				technicalQnA[fmt.Sprintf("Q. %s", question)] = fmt.Sprintf("A. %s", evaluation.Answers.TechnicalAnswers[idx])
			}
			for idx, question := range evaluation.Assignment.CodeAnalysisQuestions {
				codeQnA[fmt.Sprintf("Q. %s", question)] = fmt.Sprintf("A. %s", evaluation.Answers.CodeAnalysisAnswers[idx])
			}
			// todo: add dynamic requirement
			jobRequirement := []string{"basic Go knowledge"}
			prompt := prompts.GenerateEvaluationPrompt(&prompts.EvaluationInput{
				RequirementFactors: jobRequirement,
				Code:               evaluation.Answers.Code,
				TechnicalQnA:       technicalQnA,
				CodeQnA:            codeQnA,
			})
			openAICall(prompt)
		}
	}
}

// this is one of the goroutines as part of the consumer group
func (ap *AssessmentPipeline) assignmentConsumer() {
	ch := make(chan *nats.Msg, ap.consumptionQueue)
	sub, err := ap.Nats.ChanSubscribe(ap.EvaluationTopic, ch)

	if err != nil {
		log.Println(err)
	}

	for msg := range ch {
		subm := &types.CandidateSubmission{}
		err := json.Unmarshal(msg.Data, subm)
		if err != nil {
			log.Println(err)
		}
		ap.assignments <- *subm
	}

	sub.Unsubscribe()
}

func (ap *AssessmentPipeline) Clean() {
	ap.Nats.Close()
}
