package misc

import "fmt"
import "strings"
import "bufio"
import "os"

/*
 * Returns the first non-empty, non-blank string value
 * Returns an empty string if none qualify
 */
func CoalesceStrings(strs ...string) string {
    for _, s := range strs {
        if len(s) > 0 {
           return s
        }
    }
    return ""
}

/*
 * Returns true if user answered something that starts with "y".
 * Returns false if user answers something that starts with "n".
 * Returns yesIsDefault if no answer given.
 */
func PromptYesNo(msg string, yesIsDefault bool) bool {
    for {
        if yesIsDefault {
            fmt.Print(msg + " [Yn]: ")
        } else {
            fmt.Print(msg + " [yN]: ")
        }

        // Read input
        inReader := bufio.NewReader(os.Stdin)
        text, _ := inReader.ReadString('\n')

        // Trim leading and trailing spaces, and convert to lower case
        text = strings.TrimSpace(strings.ToLower(text))
        if len(text) == 0 {
            return yesIsDefault
        } else if strings.HasPrefix(text, "y") {
            return true
        } else if strings.HasPrefix(text, "n") {
            return false
        }
        fmt.Println("Invalid value.  Must be 'y' or 'n'.  Try again.")
    }
}

