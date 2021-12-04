package repeat

const defaultRepeatTimes = 5

func Repeat(character string, times int) string {
	repeatTimes := getRepeatTimes(times)
	return repeat(character, repeatTimes)
}

func repeat(character string, repeatTimes int) string {
	repeated := ""
	for i := 0; i < repeatTimes; i++ {
		repeated += character
	}
	return repeated
}

func getRepeatTimes(times int) int {
	repeatTimes := defaultRepeatTimes
	if times > 0 {
		repeatTimes = times
	}
	return repeatTimes
}
