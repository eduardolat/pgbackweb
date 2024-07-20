package strutil

import "testing"

func TestRemoveLeadingSlash(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Remove leading slash",
			str:  "/test/",
			want: "test/",
		},
		{
			name: "Remove leading slash from empty string",
			str:  "",
			want: "",
		},
		{
			name: "Remove leading slash from string without leading slash",
			str:  "test/",
			want: "test/",
		},
		{
			name: "Remove leading slash from string with multiple leading slashes",
			str:  "//test/",
			want: "/test/",
		},
		{
			name: "With special characters",
			str:  "/test/!@#$%^&*()_+/",
			want: "test/!@#$%^&*()_+/",
		},
		{
			name: "With special characters (emojis)",
			str:  "/test/!@#$%^&*()_+😀/",
			want: "test/!@#$%^&*()_+😀/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLeadingSlash(tt.str); got != tt.want {
				t.Errorf("RemoveLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
