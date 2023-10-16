package prompts

import "fmt"

var expectedJSON = map[string]interface{}{
	"rating":            "",
	"explanation":       "",
	"correction_factor": "",
}

var ExpectedOutputFormat string = fmt.Sprintf(
	`analyse this code and give me a general rating, explain how correct it is 
	and rate the factor by which this code can be improved in the following json format:
	%v`,
	expectedJSON,
)
