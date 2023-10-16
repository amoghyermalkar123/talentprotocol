package assessmentpipeline

import (
	"fmt"
	"log"
	"talentprotocol/db"

	"github.com/nats-io/nats.go"
)

type AssessmentPipeline struct {
	processedEvaluation chan any
	Db                  *db.DB
}

// SetupAssessmentPipeline
func (ap *AssessmentPipeline) SetupAssessmentPipeline() {
	go ap.startPipe()
	for {
		select {
		case evaluation := <-ap.processedEvaluation:
			fmt.Println("to be put in db eval", evaluation)
		}
	}
}

func (ap *AssessmentPipeline) startPipe() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	topic := "evaluation-queue"

	sub, err := nc.Subscribe(topic, func(msg *nats.Msg) {})

	for {
		msg, err := sub.NextMsg(-1)
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
		fmt.Println(msg)
		ap.processedEvaluation <- msg
	}
}
