package hello

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name     string
		language string
		expected string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "It greets with 'Hello, World' by default", args: args{name: "", language: "", expected: "Hello, World"}},
		{name: "It greets with 'Hello, {name}' when passed a name", args: args{name: "Ignacio", language: "", expected: "Hello, Ignacio"}},
		{name: "It greets with 'Hola, {name}' when passed 'Spanish' language", args: args{name: "Ignacio", language: "Spanish", expected: "Hola, Ignacio"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Hello(test.args.name, test.args.language)
			if test.args.expected != got {
				t.Errorf("expected: %s, got: %s", test.args.expected, got)
			}
		})
	}
}
