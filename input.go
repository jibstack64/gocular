package gocular

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Holds re-usable values for the InputXXXX functions.
type Input struct {
	Retry bool

	// colourset
	Colours *ColourSet
}

func (in *Input) Choices(choices []string, prompt string) (int, error) {
	return InputChoices(choices, prompt, in.Retry, in.Colours)
}

func (in *Input) Prompt(prompt string) string {
	return InputPrompt(prompt, in.Retry, in.Colours)
}

func (in *Input) Boolean(prompt string, fallback bool) bool {
	return InputBoolean(prompt, fallback, in.Colours)
}

// Creates a new input from the values given.
func NewInput(retry bool, colours *ColourSet) *Input {
	return &Input{
		Retry:   retry,
		Colours: colours,
	}
}

// Provides `choices` to the user with a `prompt`. Returns the index
// of the chosen choice. If `retry` is `true`, when the index of the
// choice is out of range, or not an integer, the user will be prompted
// again - otherwise, the `error` generated will be returned.
func InputChoices(choices []string, prompt string, retry bool, colours *ColourSet) (int, error) {
	colours.Primary.Println(prompt)
	for i, choice := range choices {
		fmt.Printf("%d. %s\n", i+1, colours.Secondary.Sprint(choice))
	}
	for {
		colours.Secondary.Print("> ")
		var input string
		inputReader := bufio.NewReader(os.Stdin)
		input, _ = inputReader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")
		index, err := strconv.Atoi(input)
		if err != nil {
			if !retry {
				return index, err
			} else {
				ClearLine()
				continue
			}
		} else {
			if index > len(choices) || index < 1 {
				if !retry {
					return -1, errors.New("index out of range")
				} else {
					ClearLine()
					continue
				}
			} else {
				return index - 1, nil
			}
		}
	}
}

// Prints the `prompt` to the console and awaits a response from the
// user. If `retry` is `true`, when a blank input is provided from
// the user, they will be prompted again until a non-blank response
// is provided.
func InputPrompt(prompt string, retry bool, colours *ColourSet) string {
	colours.Primary.Println(prompt)
	for {
		colours.Secondary.Print("> ")
		var input string
		inputReader := bufio.NewReader(os.Stdin)
		input, _ = inputReader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")
		if retry && input == "" {
			ClearLine()
			continue
		} else {
			return input
		}
	}
}

// Prints the `prompt` given and asks the user to provide 'y' or 'n'.
// `fallback` determines the default state - if the user provides a
// character or phrase that is not `y` or `n`, the boolean value returned
// will default to `fallback`.
func InputBoolean(prompt string, fallback bool, colours *ColourSet) bool {
	yn := fmt.Sprintf("(%s/%s)",
		colours.Success.Sprint("y"), colours.Error.Sprint("n"))
	if fallback {
		yn = strings.ReplaceAll(yn, "y", "Y")
	} else {
		yn = strings.ReplaceAll(yn, "n", "N")
	}
	colours.Primary.Println(prompt)
	for {
		fmt.Printf("%s: ", yn)
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input) // lower
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		} else {
			ClearLine()
			fmt.Printf("%s: %s\n", yn, func() string {
				a := "N"
				if fallback {
					a = "Y"
				}
				return a
			}())
			return fallback
		}
	}
}
