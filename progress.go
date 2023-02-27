package gocular

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Holds presets for the progress functions.
// This is useful if you re-use progress bars or cycles in your project.
type Progress struct {
	// general
	Delay time.Duration

	// cycle
	Elements string
	DotCount int

	// bar
	BarLength      int
	ShowPercentage bool

	// colourSet
	Colours *ColourSet
}

func (pp *Progress) Cycle(runner func(done *bool, err *error), text string, success string, failure string) error {
	return ProgressCycle(runner, text, success, failure, pp.Elements, pp.Delay, pp.Colours)
}

func (pp *Progress) Dots(runner func(done *bool, err *error), text string, success string, failure string) error {
	return ProgressDots(runner, text, success, failure, pp.DotCount, pp.Delay, pp.Colours)
}

func (pp *Progress) Bar(runner func(current *int, err *error), text string, success string, failure string, size int) error {
	return ProgressBar(runner, text, success, failure, size, pp.BarLength, pp.ShowPercentage, pp.Delay, pp.Colours)
}

// Creates a new `Progress` instance given the provided values.
func NewProgress(delay time.Duration, elements string, dotCount int, barLength int, showPercentage bool, colours *ColourSet) *Progress {
	return &Progress{
		Delay:          delay,
		Elements:       elements,
		DotCount:       dotCount,
		BarLength:      barLength,
		ShowPercentage: showPercentage,
		Colours:        colours,
	}
}

// Initialises and returns a default `Progress` instance.
func DefaultProgress(colourSet *ColourSet) *Progress {
	if colourSet == nil {
		colourSet = DefaultColourSet()
	}
	return NewProgress(
		time.Millisecond*500,
		"/|\\-",
		3,
		20,
		true,
		colourSet,
	)
}

// Prints `text` with a cyclic prefix (e.g. \ -> | -> / -> -).
// The cycle characters are specified through the `cycle` slice.
// Changes to the next `cycle` string every `delay` duration.
// Stops when `done` is `true`.
func ProgressCycle(runner func(done *bool, err *error), text string, success string, failure string, cycle string, delay time.Duration, colours *ColourSet) error {
	done := false
	var err error
	go runner(&done, &err)
	i := 0
	for !done && err == nil {
		colours.Secondary.Printf("%s %s\n", string(cycle[i]), colours.Primary.Sprint(text))
		time.Sleep(delay)
		i += 1
		if i == len(cycle) {
			i = 0
		}
		ClearLine()
	}
	if err != nil {
		colours.Error.Printf(failure+"\n", err.Error())
	} else {
		colours.Success.Println(success)
	}
	return err
}

// Displays `text` with a series of trailing dots.
// Adds another dot every `delay` duration.
// The maximum number of dots is specified by the `count` integer.
// Changes to finished text when `done` is equal to `true`.
func ProgressDots(runner func(done *bool, err *error), text string, success string, failure string, count int, delay time.Duration, colours *ColourSet) error {
	done := false
	var err error
	go runner(&done, &err)
	i := 0
	for !done && err == nil {
		colours.Primary.Printf("%s%s\n", text, func() string {
			d := ""
			for x := -1; x < i; x++ {
				d += "."
			}
			return d
		}())
		time.Sleep(delay)
		i += 1
		if i == count {
			i = 0
		}
		ClearLine()
	}
	if err != nil {
		colours.Error.Printf(failure+"\n", err.Error())
	} else {
		colours.Success.Println(success)
	}
	return err
}

// Receives progress from the goroutine `runner` through `current`.
// Current is parsed as a fraction of `max` and displayed in the terminal as a progress bar.
// The length of the progress bar is determined by the `bar` value.
// The progress bar will changed to a finished state when `current` is equal to `max`.
// `delay` is the update frequency.
func ProgressBar(runner func(current *int, err *error), text string, success string, failure string, max int, bar int, showPercentage bool, delay time.Duration, colours *ColourSet) error {
	current := 0
	var err error
	go runner(&current, &err)
	cells := func() string {
		s := ""
		for x := 0; x < bar; x++ {
			if ((float64(x) / float64(bar)) * float64(max)) > float64(current) {
				s += "~"
			} else {
				s += "#"
			}
		}
		return s
	}
	full := func() int {
		return int(math.Round(float64(current) / float64(max) * 100))
	}
	fmt.Println("⏳ " + colours.Primary.Sprint(text))
	for current != max && err == nil {
		if !showPercentage {
			fmt.Printf(" ⮑ [%s]\n", colours.Secondary.Sprint(cells()))
		} else {
			fmt.Printf(" ⮑ [%s] (%s)\n",
				colours.Secondary.Sprint(cells()),
				colours.Secondary.Sprintf("%d%%", full()))
		}
		time.Sleep(delay)
		ClearLine()
	}
	ClearLine()
	if err != nil {
		if strings.Contains(failure, "%s") {
			fmt.Println("❌ " + colours.Error.Sprintf(failure, err.Error()))
		} else {
			fmt.Println("❌ " + colours.Error.Sprint(failure))
		}
		if !showPercentage {
			fmt.Printf(" ⮑ [%s]\n", colours.Error.Sprint(cells()))
		} else {
			fmt.Printf(" ⮑ [%s] (%s)\n",
				colours.Error.Sprint(cells()), colours.Error.Sprintf("%d%%", full()))
		}
	} else {
		fmt.Println("✅ " + colours.Success.Sprint(success))
		if !showPercentage {
			fmt.Printf(" ⮑ [%s]\n", colours.Secondary.Sprint(cells()))
		} else {
			fmt.Printf(" ⮑ [%s] (%s)\n",
				colours.Secondary.Sprint(cells()), colours.Success.Sprintf("%d%%", full()))
		}
	}
	return err
}
