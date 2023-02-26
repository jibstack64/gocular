package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jibstack64/gocular"
)

func main() {

	colourset := gocular.DefaultColourSet()

	sleeper := func(done *bool, err *error) {
		time.Sleep(time.Second * 5)
		*done = true
	}

	progress := gocular.DefaultProgress(colourset)

	progress.Cycle(sleeper, "downloading", "downloaded nothing.", "failed to download.")

	fmt.Println()

	progress.Dots(sleeper, "downloading", "downloaded nothing.", "failed to download.")

	fmt.Println()

	progress.Bar(func(current *int, err *error) {
		for i := 0; i < 16; i++ {
			*current = i
			if i == 13 {
				*err = errors.New("oops")
			}
			time.Sleep(progress.Delay)
		}
	}, "downloading", "downloaded nothing.", "failed to download: %s", 15)

	fmt.Println()

	fmt.Println(gocular.InputPrompt("what's your name?", true, colourset))

	fmt.Println()

	enjoyed := gocular.InputBoolean("have you enjoyed today?", false, colourset)
	if enjoyed {
		fmt.Println("\nnice!")
	} else {
		fmt.Println("\nawe, that's too bad!")
	}

	fmt.Println()

	choices := []string{
		"discord", "instagram", "reddit",
	}
	index, err := gocular.InputChoices(choices, "what is your favourite social media?", false, colourset)
	if err != nil {
		colourset.Error.Println(err.Error())
	} else {
		colourset.Primary.Println(choices[index])
	}
}
