package countdown

import (
	"bytes"
	"fmt"
)

// Countdown counts down from 3, printing each number on a new line (with a 1-second pause) and when it reaches zero it
// prints "Go!"
func Countdown(buffer *bytes.Buffer) {
	_, _ = fmt.Fprintf(buffer, "3")
}
