package repeat

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		character string
		times     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "repeats 'a' five times", args: args{character: "a"}, want: "aaaaa"},
		{name: "repeats 'b' five times", args: args{character: "b"}, want: "bbbbb"},
		{name: "repeats 'c' five times", args: args{character: "c"}, want: "ccccc"},
		{name: "repeats a character a custom number of times", args: args{character: "z", times: 9}, want: "zzzzzzzzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			character := tt.args.character
			times := tt.args.times

			if got := Repeat(character, times); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
