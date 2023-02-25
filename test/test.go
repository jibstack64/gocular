package main

import (
	"time"

	"github.com/jibstack64/gocular"
)

func main() {

	sleeper := func(done *bool) {
		time.Sleep(time.Second * 5)
		*done = true
	}

	progress := gocular.NewProgress()

	progress.Cycle(sleeper, "downloading", "downloaded nothing.")

	progress.Dots(sleeper, "downloading", "downloaded nothing.")

	progress.Bar(func(current *int) {
		for i := 0; i < 5; i++ {
			*current = i
			time.Sleep(progress.Delay)
		}
	}, "downloading", "downloaded nothing.", 4)
}
