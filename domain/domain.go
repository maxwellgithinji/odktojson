package domain

// NormalizedResponse is the response to the question that has been formatted in a human readable format
// It encapsulates the original response and the questions and/or choices that have been selected
type NormalizedResponse struct {
	QuestionType string
	Question     string
	Answer       string
	Sequence     int
}
