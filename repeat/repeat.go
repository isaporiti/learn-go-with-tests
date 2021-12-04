package repeat

const repeatTimes = 5

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < repeatTimes; i++ {
		repeated += character
	}
	return repeated
}
