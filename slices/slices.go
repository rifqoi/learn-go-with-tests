package slices

func Sum(arrs []int) int {
	sum := 0
	for _, v := range arrs {
		sum = sum + v
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	sums := []int{}
	for _, arr := range arrs {
		sums = append(sums, Sum(arr))
	}
	return sums
}
