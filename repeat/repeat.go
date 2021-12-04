package repeat

const defaultRepeatTimes = 5

func Repeat(character string, times int) string {
	repeatTimes := defaultRepeatTimes
	if times > 0 {
		repeatTimes = times
	}
	repeated := ""
	for i := 0; i < repeatTimes; i++ {
		repeated += character
	}
	return repeated
}
