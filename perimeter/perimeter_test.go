package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{
		Width:  6.0,
		Height: 12.0,
	}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
