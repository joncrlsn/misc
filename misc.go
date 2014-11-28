package misc

import "fmt"
import "strings"
import "bufio"
import "os"
import "log"
import "code.google.com/p/go.crypto/ssh/terminal"


// Prompt prompts user for input and returns their response
func Prompt(prompt string) string {
    fmt.Print(prompt)
    inReader := bufio.NewReader(os.Stdin)
    text, _ := inReader.ReadString('\n')
    return text
}

// PromptYesNo returns true if user answered something that starts with "y".
// PromptYesNo returns false if user answers something that starts with "n".
// PromptYesNo returns yesIsDefault if no answer given.
func PromptYesNo(prompt string, yesIsDefault bool) bool {
    for {
        if yesIsDefault {
            fmt.Print(prompt + " [Yn]? ")
        } else {
            fmt.Print(prompt + " [yN]? ")
        }

        // Read input
        inReader := bufio.NewReader(os.Stdin)
        text, err := inReader.ReadString('\n')
		check(err, "reading stdin")

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


// ChooseOne prompts the user to enter one of n values
// Example: answer := misc.PromptOne("Continue, Quit, or Redo? [cqr]: ", ["c", "q", "r"])
func ChooseOne(prompt string, values ...string) string {
    for {
        fmt.Print(prompt)

        // Read input
        inReader := bufio.NewReader(os.Stdin)
        text, _ := inReader.ReadString('\n')

        // Trim leading and trailing spaces, and convert to lower case
        text = strings.TrimSpace(text)
        if len(text) == 0 {
            return values[0]
        }

        for _, v := range values {
            if text == v {
                return v
            }
        }

        fmt.Println("Invalid entry, try again.")
    }
}

// PromptPassword prompts user for a password that is never echoed back to the screen.
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

// ContainsString returns true if the given slice contains the given string
func ContainsString(slice []string, value string) bool {
    for _, e := range slice { 
        if value == e { return true } 
    }
    return false
}

// CompareString returns 0 if the strings match, -1 if the first string is less than the second string, and +1 if greater
func CompareString(s1, s2 string) int {
    if s1 == s2 { return 0
    } else if s1 < s2 { return -1 }  
    return 1 
}

// ColesceStrings returns the first non-empty string value
// or it returns an empty string if none qualify
func CoalesceStrings(strs ...string) string {
    for _, s := range strs {
        if len(s) > 0 {
           return s
        }
    }
    return ""
}

// =================
// Private functions
// =================
func check(err error, action string) {
    if err != nil {
        log.Fatalf("Error %s: %v\n", action, err)
    }
}
