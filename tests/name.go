//go:build ignore

package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jibstack64/gocular"
)

var (
	/*colourSet = gocular.NewColourSet(
		colour.New(colour.FgHiYellow),
		colour.New(colour.FgHiGreen),
		colour.New(colour.FgHiBlue),
		colour.New(colour.FgHiRed),
		colour.New(colour.FgHiBlack),
	)*/
	colourSet = gocular.DefaultColourSet()
	progress  = gocular.DefaultProgress(colourSet)
	input     = gocular.NewInput(true, colourSet)
)

func processing(text string) {
	progress.Dots(func(done *bool, err *error) {
		for i := 0; i < 8; i++ {
			if i == 7 {
				*done = true
			}
			time.Sleep(time.Millisecond * 500)
		}
	}, "Processing", text, "Failure: %s!")
	time.Sleep(time.Second)
}

func main() {

	name := input.Prompt("What's your name?")

	fmt.Println()

	processing(fmt.Sprintf("%s? That's a nice name.\n", name))

	genders := []string{
		"Male", "Female", "Non-binary",
	}
	gender, _ := input.Choices(genders, "What's your gender?")

	fmt.Println()

	var s string
	switch gender {
	case 0:
		s = "Good sir!"
	case 1:
		s = "M'lady."
	case 2:
		s = "Person!"
	}
	processing(s)

	colourSet.Success.Printf("\nOkay, %s of the %s gender:\n\n", name, strings.ToLower(genders[gender]))
	day := input.Boolean("Have you had a good day?", true)
	fmt.Println()
	if day {
		processing("That's good! I'll load up your daily dose of negativity for you.")
		fmt.Println()
		progress.Bar(func(current *int, err *error) {
			for i := 0; i < 8; i++ {
				*current = i
				time.Sleep(time.Millisecond * 500)
			}
		}, "Loading negativity.", "Loaded negativity!", "Failed to load.", 7)
		fmt.Println()
	} else {
		processing("Oh! Alright, I'll load up some positivity for you!")
		fmt.Println()
		progress.Bar(func(current *int, err *error) {
			for i := 0; i < 8; i++ {
				*current = i
				if i == 6 {
					*err = errors.New("")
				}
				time.Sleep(time.Millisecond * 500)
			}
		}, "Loading positivity.", "Loaded positivity!", "Failed to load positivity.", 7)
		fmt.Println()
		processing("Oops! It seemed there was an error, one second!")
		fmt.Println()
		progress.Dots(func(done *bool, err *error) {
			for i := 0; i < 5; i++ {
				if i == 4 {
					*done = true
				}
				time.Sleep(time.Millisecond * 500)
			}
		}, "I'll try again", "Success!", "Failure: %s!")
		colourSet.Success.Println("\nHave a wonderful day!")
	}
}
