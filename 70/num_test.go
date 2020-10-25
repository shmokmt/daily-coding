package num

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{1, 19},
		{2, 28},
	}
	for _, tt := range testCases {
		if Perfect(tt.input) != tt.output {
			t.Errorf("want: %d, got: %d", tt.output, tt.input)
		}
	}
}
