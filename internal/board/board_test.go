package board

import (
	"testing"
)

func TestGenerateBoard(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected string
	}{
		{
			name:     "Размер 1",
			size:     1,
			expected: " \n",
		},
		{
			name:     "Размер 2",
			size:     2,
			expected: " #\n# \n",
		},
		{
			name:     "Размер 3",
			size:     3,
			expected: " # \n# #\n # \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateBoard(tt.size)
			if result != tt.expected {
				t.Errorf("GenerateBoard(%d) = %q, want %q", tt.size, result, tt.expected)
			}
		})
	}
}
