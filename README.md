misc
====

Miscellaneous GoLang utility methods I've written and collected.

/*
 * Returns true if user answered something that starts with "y".
 * Returns false if user answers something that starts with "n".
 * Returns yesIsDefault if no answer given.
 */
func PromptYesNo(msg string, yesIsDefault bool) bool

/*
 * Returns the first non-empty, non-blank string value
 * Returns an empty string if none qualify
 */
func CoalesceStrings(strs ...string) string
