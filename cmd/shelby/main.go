package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "shelby",
		Usage: "Help shelby learn!",
		Action: func(c *cli.Context) error {
			fmt.Printf("Hi, I'm Shelby, my memory isn't so good...")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
