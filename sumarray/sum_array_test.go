package sumarray

import "testing"

func TestSumArray(t *testing.T) {
	type args struct {
		numbers []int
		want    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "it sums the numbers in [2, 2]", args: args{numbers: []int{2, 2}, want: 4}},
		{name: "it sums the numbers in [1, 2, 3, 4, 5]", args: args{numbers: []int{1, 2, 3, 4, 5}, want: 15}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Sum(test.args.numbers); got != test.args.want {
				t.Errorf("Sum(numbers): got %d, want %d", got, test.args.want)
			}
		})
	}
}

func Sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
