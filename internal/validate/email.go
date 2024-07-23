package validate

import "regexp"

// Email validates an email address.
// It returns a boolean indicating whether
// the email is valid or not.
func Email(email string) bool {
	// Regular expression to match email format
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(regex)

	// Match the email against the regular expression
	return re.MatchString(email)
}
