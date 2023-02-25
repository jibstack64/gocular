package gocular

import (
	"math"
	"time"
)

const (
	DEFAULT_DELAY           = time.Millisecond * 500
	DEFAULT_ELEMENTS        = "/|\\-"
	DEFAULT_DOT_COUNT       = 3
	DEFAULT_BAR_LENGTH      = 20
	DEFAULT_SHOW_PERCENTAGE = true
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

	// colourset
	Colours *ColourSet
}

func (pp *Progress) Cycle(runner func(done *bool), text string, finished string) {
	ProgressCycle(runner, text, finished, pp.Elements, pp.Delay, pp.Colours)
}

func (pp *Progress) Dots(runner func(done *bool), text string, finished string) {
	ProgressDots(runner, text, finished, pp.DotCount, pp.Delay, pp.Colours)
}

func (pp *Progress) Bar(runner func(current *int), text string, finished string, size int) {
	ProgressBar(runner, text, finished, size, pp.BarLength, pp.ShowPercentage, pp.Delay, pp.Colours)
}

// Initialises and returns a default `Progress` instance.
func NewProgress() *Progress {
	return &Progress{
		Delay:          DEFAULT_DELAY,
		Elements:       DEFAULT_ELEMENTS,
		DotCount:       DEFAULT_DOT_COUNT,
		BarLength:      DEFAULT_BAR_LENGTH,
		ShowPercentage: DEFAULT_SHOW_PERCENTAGE,
		Colours:        NewColourSet(),
	}
}

// Prints `text` with a cyclic prefix (e.g. \ -> | -> / -> -).
// The cycle characters are specified through the `cycle` slice.
// Changes to the next `cycle` string every `delay` duration.
// Stops when `done` is `true`.
func ProgressCycle(runner func(done *bool), text string, finished string, cycle string, delay time.Duration, colours *ColourSet) {
	done := false
	go runner(&done)
	i := 0
	for !done {
		colours.Secondary.Printf("%s %s\n", string(cycle[i]), colours.Primary.Sprint(text))
		time.Sleep(delay)
		i += 1
		if i == len(cycle) {
			i = 0
		}
		ClearLine()
	}
	colours.Success.Println(finished)
}

// Displays `text` with a series of trailing dots.
// Adds another dot every `delay` duration.
// The maximum number of dots is specified by the `count` integer.
// Changes to finished text when `done` is equal to `true`.
func ProgressDots(runner func(done *bool), text string, finished string, count int, delay time.Duration, colours *ColourSet) {
	done := false
	go runner(&done)
	i := 0
	for !done {
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
	colours.Success.Println(finished)
}

// Receives progress from the goroutine `runner` through `current`.
// Current is parsed as a fraction of `max` and displayed in the terminal as a progress bar.
// The length of the progress bar is determined by the `bar` value.
// The progress bar will changed to a finished state when `current` is equal to `max`.
// `delay` is the update frequency.
func ProgressBar(runner func(current *int), text string, finished string, max int, bar int, showPercentage bool, delay time.Duration, colours *ColourSet) {
	current := 0
	go runner(&current)
	for current != max {
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
		}()
		if !showPercentage {
			colours.Bracket.Printf("%s [%s]\n", colours.Primary.Sprint(text), colours.Secondary.Sprint(cells))
		} else {
			colours.Bracket.Printf("%s [%s] (%s)\n",
				colours.Primary.Sprint(text), colours.Secondary.Sprint(cells),
				colours.Secondary.Sprintf("%d%%", int(math.Round(float64(current)/float64(max)*100))))
		}
		time.Sleep(delay)
		ClearLine()
	}
	colours.Success.Println(finished)
}
