misc
====

Miscellaneous GoLang utility methods I've written and collected from various public sources on the web.

```
/*
 * Returns true if user answered something that starts with "y".
 * Returns false if user answers something that starts with "n".
 * Returns yesIsDefault if no answer given.
 */
func PromptYesNo(msg string, yesIsDefault bool) bool
```

```
/*
 * Prompts user for input and returns their response
 */
func Prompt(msg string) string 
```

```
/*
 * Prompts user for a password that is never echoed back to the screen.
 */
func PromptPassword(prompt string) string
```

```
/*
 * Returns the first non-empty, non-blank string value
 * Returns an empty string if none qualify
 */
func CoalesceStrings(strs ...string) string
```
