## Input

### `Input` : `struct`

Holds re-usable values for the input functions.

  - ### `Choices(...)`, `Prompt(...)` and `Boolean(...)`

    These functions are relays for the `InputChoices`, `InputPrompt` and `InputBoolean` functions.

### `InputChoices(choices []string, ...)`

Presents each of the `choices` to the user with the `prompt`. When the user picks a choice by number, the choice's index in the `choices` slice is returned. If `retry` is `true`, when the choice is out of range or not a valid integer, the error will be returned rather than be ignored and prompt the user again.

> ![Preview](https://github.com/jibstack64/gocular/blob/master/examples/choices.png)

### `InputPrompt(prompt string, ...)`

Prompts the user with `prompt` and expects a response. If no response is provided and `retry` is `true`, the user will be prompted again. If `retry` is `false`, the empty response is returned.

> ![Preview](https://github.com/jibstack64/gocular/blob/master/examples/prompt.png)

### `InputBoolean(prompt string, ...)`

Prints the `prompt` given and asks the user to provide `y` or `n` in traditional UNIX style. If `fallback` is `true`, then `y` will be the default option and will be capitalised in the console, if `false`: the same but for `n`.

> ![Preview](https://github.com/jibstack64/gocular/blob/master/examples/boolean.png)
