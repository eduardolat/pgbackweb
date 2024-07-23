package validate

import "testing"

func TestEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"", false},
		{"test", false},
		{"test@", false},
		{"@example.com", false},
		{"test@example", false},
		{"test@example.com", true},
		{"test@example.com.gt", true},
	}

	for _, testItem := range tests {
		isValid := Email(testItem.email)
		if isValid != testItem.valid {
			t.Errorf("Email(%s) expected %v, got %v", testItem.email, testItem.valid, isValid)
		}
	}
}
