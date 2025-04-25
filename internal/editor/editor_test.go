package editor_test

import (
	"testing"

	detector "github.com/tahmazidik/Copy_Paste/internal/editor"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Empty string", "", false},
		{"Valid string", "Hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detector.Check(tt.input); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
