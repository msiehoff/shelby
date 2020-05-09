package shelby

import (
	"fmt"
	"io"
)

// Challenge - the smallest unit of spaced repetition
// could be anything from a question/answer to a coding
// problem
type Challenge interface {
	// present the challenge to the user
	Present(w io.Writer) error

	// ShowAnswer - display the correct answer
	ShowAnswer(w io.Writer) error

	// let the user input how difficult the challenge was
	// to inform when they'll see it next
	//Respond(difficulty int) error
}

// ChallengeLoader - Loads challenges from a source
type ChallengeLoader interface {
	// Load challenges from a given source
	Load() ([]Challenge, error)
}

// SimpleChallengeLoader -
type SimpleChallengeLoader struct {
	Challenges []Challenge
}

// Load -return Challenges
func (s SimpleChallengeLoader) Load() ([]Challenge, error) {
	return s.Challenges, nil
}

// Flashcard  - a simple question & answer challenge
type Flashcard struct {
	Answer   string
	Question string
}

// Present - display the question
func (f Flashcard) Present(w io.Writer) error {
	qText := fmt.Sprintf("\n\n%s\n\n", f.Question)
	fmt.Fprintf(w, promptColor, qText)
	return nil
}

// ShowAnswer - show the answer
func (f Flashcard) ShowAnswer(w io.Writer) error {
	fmt.Fprintf(w, noticeColor, "\n------- Answer -----------------------------")
	_, err := fmt.Fprintf(w, "\n\n%s\n\n", f.Answer)
	fmt.Fprintf(w, noticeColor, "--------------------------------------------\n\n")

	return err
}
