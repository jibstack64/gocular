## Progress

### `Progress` : `struct`

Holds presets for the progress functions. This is useful for re-using progress bars and cycles.

  - ### `Cycle(...)`, `Dots(...)` and `Bar(...)`

    Relays for the `ProgressCycle`, `ProgressDots` and `ProgressBar` functions that fill many of the tedious arguments with the values stored within `Progress`.

### `NewProgress(delay time.Duration, ...)`

Creates a new instance of `Progress` with the given values.

### `DefaultProgress(colourSet *ColourSet)`

Initialises and returns a default `Progress` object. If `colourSet` is `nil`, a `DefaultColourSet()` is generated instead.

### `ProgressCycle(...)`, `ProgressDots(...)` and `ProgressBar(...)`

All of these functions take tedious and often obsolete arguments for the sake of customability. To simplify the use of these functions, use a `Progress` object - it saves time, data and is easier to understand. Further documentation can be found within the source code, however I *heavily* suggest you use a `Progress` struct.

> ![Preview](https://github.com/jibstack64/gocular/blob/master/examples/cycle_dots_bar.gif)
