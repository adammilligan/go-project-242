package format

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		{
			name:     "bytes",
			size:     123,
			human:    false,
			expected: "123B",
		},
		{
			name:     "bytes-big",
			size:     25165824,
			human:    false,
			expected: "25165824B",
		},
		{
			name:     "human-mb",
			size:     25165824,
			human:    true,
			expected: "24.0MB",
		},
		{
			name:     "human-mb-smaller",
			size:     1234567,
			human:    true,
			expected: "1.2MB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, FormatSize(tc.size, tc.human))
		})
	}
}
