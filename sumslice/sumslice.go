package sumslice

func Sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	var sliceSum int
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sliceSum = Sum(tail)
		} else {
			sliceSum = 0
		}
		sums = append(sums, sliceSum)
	}
	return sums
}
