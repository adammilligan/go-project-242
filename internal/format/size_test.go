package format

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatSize(t *testing.T) {
	t.Run("invariant scenarios", func(t *testing.T) {
		t.Run("panics on negative size", func(t *testing.T) {
			require.Panics(t, func() {
				FormatSize(-1, false)
			})
		})
	})

	type testCase struct {
		name     string
		size     int64
		human    bool
		expected string
	}

	run := func(t *testing.T, tc testCase) {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, FormatSize(tc.size, tc.human))
		})
	}

	t.Run("success scenarios", func(t *testing.T) {
		tests := []testCase{
			{
				name:     "human=false returns bytes without conversion",
				size:     123,
				human:    false,
				expected: "123B",
			},
			{
				name:     "human=false keeps large sizes in bytes",
				size:     25165824,
				human:    false,
				expected: "25165824B",
			},
			{
				name:     "human=true converts bytes to MB",
				size:     25165824,
				human:    true,
				expected: "24.0MB",
			},
			{
				name:     "human=true converts bytes to MB with rounding",
				size:     1234567,
				human:    true,
				expected: "1.2MB",
			},
		}

		for _, tc := range tests {
			run(t, tc)
		}
	})
}
