package strutil

import "testing"

func TestRemoveTrailingSlash(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Remove trailing slash",
			str:  "/test/",
			want: "/test",
		},
		{
			name: "Remove trailing slash from empty string",
			str:  "",
			want: "",
		},
		{
			name: "Remove trailing slash from string without trailing slash",
			str:  "/test",
			want: "/test",
		},
		{
			name: "Remove trailing slash from string with multiple trailing slashes",
			str:  "/test//",
			want: "/test/",
		},
		{
			name: "With special characters",
			str:  "/test/!@#$%^&*()_+/",
			want: "/test/!@#$%^&*()_+",
		},
		{
			name: "With special characters (emojis)",
			str:  "/test/!@#$%^&*()_+😀/",
			want: "/test/!@#$%^&*()_+😀",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveTrailingSlash(tt.str); got != tt.want {
				t.Errorf("RemoveTrailingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
