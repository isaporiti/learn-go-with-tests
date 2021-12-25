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
	for _, numbers := range numbersToSum {
		sums = append(sums, sumSlice(numbers))
	}
	return sums
}

func sumSlice(numbers []int) int {
	if empty(numbers) {
		return 0
	}
	return sumTail(numbers)
}

func empty(numbers []int) bool {
	return len(numbers) == 0
}

func sumTail(numbers []int) int {
	tail := numbers[1:]
	return Sum(tail)
}
