## Colour

### `ColourSet` : `struct`

Holds `Primary`, `Secondary`, `Success`, `Error` and `Bracket` colours. These are all used at various points in the program.

  - ### `Update(primary *colour.Color, ...)`
    
    Used to update the values of a `ColourSet` object.

### `NewColourSet(primary *colour.Color, ...)`

Returns a newly created `ColourSet` object from the colours provided. This can be used to create your own "themes" instead of using the default.

### `DefaultColourSet()`

Generates the default `ColourSet` object: light cyan as primary, light magenta as secondary, green as success, red as error and grey as bracket.
