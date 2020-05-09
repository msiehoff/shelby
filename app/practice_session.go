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

// question header
// questioon
// answer
// info

// Practice -
func (p *Practicer) Practice() error {
	challenges, err := p.Loader.Load()
	if err != nil {
		return err
	}
	p.challenges = challenges

	scanner := bufio.NewScanner(p.Reader)
	challCount := 0
	for {
		nextChallenge, hasMoreChallenges := p.Next()
		if !hasMoreChallenges {
			break
		}

		// TODO: pass io.Writer to challenge
		challCount++
		p.questionHeader(challCount, len(challenges))
		nextChallenge.Present()
		p.postChallengeText()

		// enter anything to view answer
		// if yes: how difficult was it? (10 if you didn't get it)
		for scanner.Scan() {
			scanner.Text()
			nextChallenge.ShowAnswer(p.Writer)
			p.askDifficulty()
			break
		}

		for scanner.Scan() {
			// difficulty answer
			msg := scanner.Text()

			diff, err := strconv.Atoi(msg)
			if err != nil {
				p.invalidDiffMsg(msg)
				continue
			}

			if diff < 1 || diff > 10 {
				p.invalidDiffMsg(msg)
				continue
			}

			// send feedback to question
			break
		}
	}

	fmt.Printf("\nGreat Job, you're done!!!\n")

	return nil
}

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
)

func (p *Practicer) questionHeader(ind, count int) {
	qText := fmt.Sprintf("------------ %d/%d------------", ind, count)
	fmt.Fprintf(p.Writer, infoColor, qText)
}

func (p *Practicer) postChallengeText() {
	txt := "When you're ready to see the answer hit any key"
	fmt.Fprintf(p.Writer, infoColor, txt)
}

func (p *Practicer) askDifficulty() {
	txt := "How difficult was that question? (1-10)"
	fmt.Fprintf(p.Writer, infoColor, txt)
}

func (p *Practicer) invalidDiffMsg(msg string) {
	msg = fmt.Sprintf("You said: %s, please enter a valid integer 1-10", msg)
	fmt.Fprintf(p.Writer, errorColor, msg)
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
