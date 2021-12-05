package sumslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
		want    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "slice of [1, 2, 3, 4, 5]", args: args{numbers: []int{1, 2, 3, 4, 5}, want: 15}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Sum(test.args.numbers); got != test.args.want {
				t.Errorf("Sum(numbers): got %d, want %d", got, test.args.want)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 2}, []int{5, 6})
	want := []int{4, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
