package shelby

import (
	"fmt"
	"io"
)

// Challenge - TODO
type Challenge interface {
	// present the challenge to the user
	Present() error

	// ShowAnswer - display the correct answer
	ShowAnswer(w io.Writer) error

	// let the user input how difficult the challenge was
	// to inform when they'll see it next
	// TODO
	//Respond(difficulty int) error
}

// ChallengeLoader - TODO
type ChallengeLoader interface {
	// Load challenges from a given source
	Load() ([]Challenge, error)
}

// SimpleChallengeLoader -
type SimpleChallengeLoader struct {
	Challenges []Challenge
}

// Load -
func (s SimpleChallengeLoader) Load() ([]Challenge, error) {
	return s.Challenges, nil
}

// Flashcard  -
type Flashcard struct {
	Answer   string
	Question string
}

// Present -
func (f Flashcard) Present() error {
	fmt.Printf("\n\n%s\n\n", f.Question)
	return nil
}

// ShowAnswer -
func (f Flashcard) ShowAnswer(w io.Writer) error {
	fmt.Fprintf(w, noticeColor, "\n------- Answer -----------------------------")
	_, err := fmt.Fprintf(w, "\n\n%s\n\n", f.Answer)
	fmt.Fprintf(w, noticeColor, "--------------------------------------------\n\n")

	return err
}
