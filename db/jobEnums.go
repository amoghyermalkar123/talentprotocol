package db

// Define a type for your "enum"
type Status int

const (
	Open Status = iota
	Closed
	UnderEvaluation
	Accepted
)

func (s Status) String() string {
	return [...]string{"Open", "Closed", "Under Evaluation", "Accepted"}[s]
}
