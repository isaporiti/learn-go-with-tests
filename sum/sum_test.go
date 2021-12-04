package sum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		x, y     int
		expected int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "it sums 2 and 2", args: args{x: 2, y: 2, expected: 4}},
		{name: "it sums 9 and 1", args: args{x: 9, y: 1, expected: 10}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			x := test.args.x
			y := test.args.y

			got := Sum(x, y)

			expected := test.args.expected
			if expected != got {
				t.Errorf("expected: %d, got: %d", expected, got)
			}
		})
	}
}
