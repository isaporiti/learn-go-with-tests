package repeat

import "testing"

func TestRepeat1(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "repeats 'a' five times", args: args{character: "a"}, want: "aaaaa"},
		{name: "repeats 'b' five times", args: args{character: "b"}, want: "bbbbb"},
		{name: "repeats 'c' five times", args: args{character: "c"}, want: "ccccc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.character); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
