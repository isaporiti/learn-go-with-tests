package sumslice

func Sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))
	for i := range numbers {
		for j := range numbers[i] {
			sums[i] += numbers[i][j]
		}
	}
	return sums
}
