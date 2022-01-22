package countdown

import (
	"fmt"
	"io"
)

// Countdown counts down from 3, printing each number on a new line (with a 1-second pause) and when it reaches zero it
// prints "Go!"
func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		_, _ = fmt.Fprintln(writer, i)
	}
	_, _ = fmt.Fprint(writer, "Go!")
}
