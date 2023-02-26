package gocular

import colour "github.com/fatih/color"

// Holds primary, secondary, success, error and bracket colours.
type ColourSet struct {
	Primary   *colour.Color
	Secondary *colour.Color
	Success   *colour.Color
	Error     *colour.Color
	Bracket   *colour.Color
}

// Updates the `ColourSet`'s values.
func (colours *ColourSet) Update(primary *colour.Color, secondary *colour.Color, success *colour.Color, error_ *colour.Color, bracket *colour.Color) {
	if primary != nil {
		colours.Primary = primary
	}
	if secondary != nil {
		colours.Secondary = secondary
	}
	if success != nil {
		colours.Success = success
	}
	if error_ != nil {
		colours.Error = error_
	}
	if bracket != nil {
		colours.Bracket = bracket
	}
}

// Returns a new `ColourSet` from the provided values.
func NewColourSet(primary *colour.Color, secondary *colour.Color, success *colour.Color, error_ *colour.Color, bracket *colour.Color) *ColourSet {
	return &ColourSet{
		Primary:   primary,
		Secondary: secondary,
		Success:   success,
		Error:     error_,
		Bracket:   bracket,
	}
}

// Creates and returns a default `ColourSet` object.
func DefaultColourSet() *ColourSet {
	return NewColourSet(
		colour.New(colour.FgHiCyan),
		colour.New(colour.FgHiMagenta),
		colour.New(colour.FgHiGreen),
		colour.New(colour.FgHiRed),
		colour.New(colour.FgHiBlack),
	)
}
