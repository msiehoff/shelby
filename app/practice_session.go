package shelby

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

		// TODO: pass io.Writer to challenge
		nextChallenge.Present()

		// enter anything to view answer
		// if yes: how difficult was it? (10 if you didn't get it)
		//scanner.Scan()
		//scanner.Text()
		//nextChallenge.ShowAnswer(p.Writer)

		for scanner.Scan() {
			// user answered question
			msg := scanner.Text()

			diff, err := strconv.Atoi(msg)
			if err != nil {
				msg = invalidDiffMsg(msg)
				fmt.Fprint(p.Writer, msg)
				continue
			}

			if diff < 1 || diff > 10 {
				msg = invalidDiffMsg(msg)
				fmt.Fprint(p.Writer, msg)
				continue
			}

			// send feedback to question
			break
		}
	}

	fmt.Printf("\nGreat Job, you're done!!!\n")

	return nil
}

func invalidDiffMsg(msg string) string {
	return fmt.Sprintf("\n\nYou said: %s, please enter a valid integer 1-10\n\n", msg)
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
