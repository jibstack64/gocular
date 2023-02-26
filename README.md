# gocular

![GitHub](https://img.shields.io/github/license/jibstack64/gocular) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jibstack64/gocular) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/jibstack64/gocular)

*A simple TUI library for Go.*

![Preview](https://github.com/jibstack64/gocular/blob/master/preview.gif)

---

## Colour

### `ColourSet` : `struct`

Holds `Primary`, `Secondary`, `Success`, `Error` and `Bracket` colours. These are all used at various points in the program.

  - ### `Update(primary *colour.Color, ...)`
    
    Used to update the values of a `ColourSet` object.

### `NewColourSet(primary *colour.Color, ...)`

Returns a newly created `ColourSet` object from the colours provided. This can be used to create your own "themes" instead of using the default.

### `DefaultColourSet()`

Generates the default `ColourSet` object: light cyan as primary, light magenta as secondary, green as success, red as error and grey as bracket.

---

## Progress

### `Progress` : `struct`

Holds presets for the progress functions. This is useful for re-using progress bars and cycles.

  - ### `Cycle(...)`, `Dots(...)` and `Bar(...)`

    Relays for the `ProgressCycle`, `ProgressDots` and `ProgressBar` functions that fill many of the tedious arguments with the values stored within `Progress`.

### `NewProgress(delay time.Duration, ...)`

Creates a new instance of `Progress` with the given values.

### `DefaultProgress(colourSet *ColourSet)

Initialises and returns a default `Progress` object. If `colourSet` is `nil`, a `DefaultColourSet()` is generated instead.

### `ProgressCycle(...)`, `ProgressDots(...)` and `ProgressBar(...)`

All of these functions take tedious and often obsolete arguments for the sake of customability. To simplify the use of these functions, use a `Progress` object - it saves time, data and is easier to understand. Further documentation can be found within the source code, however I *heavily* suggest you use a `Progress` struct.

---

## Input

### `Input` : `struct`

Holds re-usable values for the input functions.

  - ### `Choices(...)`, `Prompt(...)` and `Boolean(...)`

    These functions are relays for the `InputChoices`, `InputPrompt` and `InputBoolean` functions.

### `InputChoices(choices []string, ...)`

Presents each of the `choices` to the user with the `prompt`. When the user picks a choice by number, the choice's index in the `choices` slice is returned. If `retry` is `true`, when the choice is out of range or not a valid integer, the error will be returned rather than be ignored and prompt the user again.

### `InputPrompt(prompt string, ...)`

Prompts the user with `prompt` and expects a response. If no response is provided and `retry` is `true`, the user will be prompted again. If `retry` is `false`, the empty response is returned.

### `InputBoolean(prompt string, ...)`

Prints the `prompt` given and asks the user to provide `y` or `n` in traditional UNIX style. If `fallback` is `true`, then `y` will be the default option and will be capitalised in the console, if `false`: the same but for `n`.

---

## Terminal

### `ClearLine()`

Clears the previous line from the console.

### `ClearLines(x int)`

Clears `x` previous lines from the console.
