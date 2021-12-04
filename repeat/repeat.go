package repeat

func Repeat(character string, times int) string {
	if times == 0 {
		return ""
	}
	return repeat(character, times)
}

func repeat(character string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
