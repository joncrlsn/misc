package misc

import "fmt"
import "strings"
import "bufio"
import "os"
import "log"
import "code.google.com/p/go.crypto/ssh/terminal"

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
 * Prompts user for input and returns their response
 */
func Prompt(prompt string) string {
    fmt.Print(prompt)
    inReader := bufio.NewReader(os.Stdin)
    text, _ := inReader.ReadString('\n')
    return text
}

/*
 * Returns true if user answered something that starts with "y".
 * Returns false if user answers something that starts with "n".
 * Returns yesIsDefault if no answer given.
 */
func PromptYesNo(prompt string, yesIsDefault bool) bool {
    for {
        if yesIsDefault {
            fmt.Print(prompt + " [Yn]? ")
        } else {
            fmt.Print(prompt + " [yN]? ")
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

/*
 * Prompts user for a password that is never echoed back to the screen.
 */
func PromptPassword(prompt string) string {
    stdin := 1
    fmt.Fprintf(os.Stdin, prompt)

    // Get current state of terminal
//    s, err := terminal.MakeRaw(stdin)
//    check(err, "making raw terminal, Saving old terminal state")
//    defer terminal.Restore(stdin, s)

    // Read password from stdin
    b, err := terminal.ReadPassword(stdin)
    check(err, "reading from terminal")

    return string(b)
}

func check(err error, action string) {
    if err != nil {
        log.Fatalf("Error %s: %v\n", action, err)
    }
}
