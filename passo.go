package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

var num = 6

const alphaset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const mixedset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithAlphaset(length int, alphaset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphaset[seededRand.Intn(len(alphaset))]
	}
	return string(b)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringWithMixedset(length int, mixedset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = mixedset[seededRand.Intn(len(mixedset))]
	}
	return string(b)
}

func main() {
	app := &cli.App{
		Name:  "passo",
		Usage: "Simple Password generator CLI",
		Authors: []*cli.Author{
			{
				Name:  "uzxrk",
				Email: "uzair558@gmail.com",
			},
		},
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "alpha",
				Aliases: []string{"a"},
				Usage:   "Generate a random alphabetic password.",
				Subcommands: []*cli.Command{
					{
						Name:  strconv.Itoa(num),
						Usage: fmt.Sprintf("Generate a password with length of %d characters.", num),
						Action: func(cCtx *cli.Context) error {
							StringWithAlphaset(num, alphaset)
							return nil
						},
					},
				},
			},
			{
				Name:    "alphanum",
				Aliases: []string{"an"},
				Usage:   "Generate a random alpha-numeric password.",
				Subcommands: []*cli.Command{
					{
						Name:  strconv.Itoa(num),
						Usage: fmt.Sprintf("Generate a password with length of %d characters.", num),
						Action: func(cCtx *cli.Context) error {
							StringWithCharset(num, charset)
							return nil
						},
					},
				},
			},
			{
				Name:    "mixed",
				Aliases: []string{"mx"},
				Usage:   "Generate a random alpha-numeric password with special characters.",
				Subcommands: []*cli.Command{
					{
						Name:  strconv.Itoa(num),
						Usage: fmt.Sprintf("Generate a password with length of %d characters.", num),
						Action: func(c *cli.Context) error {
							StringWithMixedset(num, mixedset)
							return nil
						},
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
