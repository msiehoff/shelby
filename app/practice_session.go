package shelby

import (
	"bufio"
	"fmt"
	"io"
)

// Practicer -
type Practicer struct {
	Loader ChallengeLoader
	Reader io.Reader
	Writer io.Writer

	challengeInd int
	challenges   []Challenge
}

// Practice -
func (p *Practicer) Practice() error {
	challenges, err := p.Loader.Load()
	if err != nil {
		return err
	}
	p.challenges = challenges

	scanner := bufio.NewScanner(p.Reader)
	for {
		nextChallenge, hasMoreChallenges := p.Next()
		if !hasMoreChallenges {
			break
		}

		nextChallenge.Present()

		p.Writer.Write([]byte(`\n\nHow difficult was that for you on a scale of 1-10?\n`))

		scanner.Scan()
		msg := scanner.Text()

		fmt.Printf("\n\nYou said: %s\n\n", msg)

		// convert to int difficulty
		// send feedback to question
	}

	fmt.Printf("\nGreat Job, you're done!!!\n")

	return nil
}

// Next -
func (p *Practicer) Next() (Challenge, bool) {
	if !p.hasMoreChallenges() {
		return nil, false
	}

	return p.nextChallenge(), true
}

func (p *Practicer) hasMoreChallenges() bool {
	return p.challengeInd < len(p.challenges)
}

func (p *Practicer) nextChallenge() Challenge {
	nextChallenge := p.challenges[p.challengeInd]
	p.challengeInd++

	return nextChallenge
}
