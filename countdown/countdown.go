package countdown

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown counts down from 3, printing each number on a new line (with a 1-second pause) and when it reaches zero it
// prints "Go!"
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	_, _ = fmt.Fprint(writer, finalWord)
}
