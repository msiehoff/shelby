package shelby_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/bmizerany/assert"
	shelby "github.com/msiehoff/shelby/app"
)

func TestPracticer_Practice(t *testing.T) {
	// presents all challenges
	// shows answer for all challenges
	// validates difficulty
	mockChallenges := []shelby.Challenge{
		&mockChallenge{},
		&mockChallenge{},
	}
	loader := shelby.SimpleChallengeLoader{
		Challenges: mockChallenges,
	}

	p := &shelby.Practicer{
		Loader: loader,
		Reader: &bytes.Reader{},
		Writer: &bytes.Buffer{},
	}

	err := p.Practice()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("PresentsAllChallenges", func(t *testing.T) {
		var numPresented int
		for _, c := range mockChallenges {
			mc, isMock := c.(*mockChallenge)
			if !isMock {
				continue
			}

			if !mc.Presented {
				continue
			}

			numPresented++
		}

		assert.Equal(t, len(mockChallenges), numPresented)
	})

}

type mockChallenge struct {
	Answered  bool
	Presented bool
}

func (m *mockChallenge) Present(w io.Writer) error {
	m.Presented = true
	return nil
}

func (m *mockChallenge) ShowAnswer(w io.Writer) error {
	m.Answered = true
	return nil
}
