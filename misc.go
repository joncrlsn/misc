package misc

import "fmt"
import "strings"
import "bufio"
import "os"
import "os/signal"
import "syscall"
import "log"
import "golang.org/x/crypto/ssh/terminal"

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
// This needs to be improved because it leaves terminal in unusable state when
// user presses Ctrl-C
func PromptPassword(prompt string) string {
	stdin := 1
	fmt.Fprintf(os.Stdin, prompt)

	// Get current state of terminal
	s, err := terminal.MakeRaw(stdin)
	check(err, "making raw terminal, Saving old terminal state")
	defer terminal.Restore(stdin, s)

	// trap Ctrl-C and restore screen
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		terminal.Restore(stdin, s)
		os.Exit(1)
	}()

	// Read password from stdin
	b, err := terminal.ReadPassword(stdin)
	check(err, "reading from terminal")

	return string(b)
}

// InStrings returns true if the first argument is in the list of following arguments
// or it returns false if none match. InStrings is similar to ConstainsString which
// takes a slice of strings instead of a variable number of string arguments.
func InStrings(str string, strs ...string) bool {
	for _, s := range strs {
		if str == s {
			return true
		}
	}
	return false
}

// ContainsString returns true if the given slice contains the given string
func ContainsString(slice []string, value string) bool {
	for _, e := range slice {
		if value == e {
			return true
		}
	}
	return false
}

// CompareStrings returns negative when first string is less, positive if more, zero otherwise
func CompareStrings(s1, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 < s2 {
		return -1
	}
	return 1
}

// ColesceStrings returns the first non-empty string value
// or, if none qualify, it returns an empty string
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
