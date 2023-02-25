package gocular

import colour "github.com/fatih/color"

const (
	DEFAULT_PRIMARY   = colour.FgHiCyan
	DEFAULT_SECONDARY = colour.FgHiMagenta
	DEFAULT_SUCCESS   = colour.FgHiGreen
	DEFAULT_ERROR     = colour.FgHiRed
	DEFAULT_BRACKET   = colour.FgHiBlack
)

// Holds primary, secondary, success, error and bracket colours.
type ColourSet struct {
	Primary   *colour.Color
	Secondary *colour.Color
	Success   *colour.Color
	Error     *colour.Color
	Bracket   *colour.Color
}

// Creates and returns a default `ColourSet` object.
func NewColourSet() *ColourSet {
	return &ColourSet{
		Primary:   colour.New(DEFAULT_PRIMARY),
		Secondary: colour.New(DEFAULT_SECONDARY),
		Success:   colour.New(DEFAULT_SUCCESS),
		Error:     colour.New(DEFAULT_ERROR),
		Bracket:   colour.New(DEFAULT_BRACKET),
	}
}
