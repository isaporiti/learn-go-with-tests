package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}

	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	testCases := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{6.0, 12.0}, want: 72.0},
		{shape: Circle{10.0}, want: 314.1592653589793},
		{shape: Triangle{12.0, 6.0}, want: 36},
	}

	for _, testCase := range testCases {
		got := testCase.shape.Area()
		want := testCase.want

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}
}
