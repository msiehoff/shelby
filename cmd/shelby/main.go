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
				QuestionText: "Answer to Question 1?",
			},
			shelby.Flashcard{
				QuestionText: "Answer to Question 2?",
			},
		},
	}

	return &shelby.Practicer{
		Loader: loader,
		Reader: bufio.NewReader(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
	}
}
