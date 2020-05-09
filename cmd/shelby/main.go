package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	shelby "github.com/msiehoff/shelby/app"
	"github.com/urfave/cli"
)

func main() {
	practicer := buildPracticer()

	fmt.Printf("\n\nHey there, ready for some challenges?\n\n")

	app := &cli.App{
		Name:  "shelby",
		Usage: "Help shelby learn!",
		Action: func(c *cli.Context) error {
			practicer.Practice()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func buildPracticer() *shelby.Practicer {
	loader := shelby.SimpleChallengeLoader{
		Challenges: []shelby.Challenge{
			shelby.Flashcard{
				Question: "How far away is the Sun?",
				Answer:   "It's 8 light minutes away.",
			},
			shelby.Flashcard{
				Question: "How far away is the moon in light seconds?",
				Answer:   "It's 1 light second away.",
			},
		},
	}

	return &shelby.Practicer{
		Loader: loader,
		Reader: bufio.NewReader(os.Stdin),
		Writer: os.Stdout,
	}
}
