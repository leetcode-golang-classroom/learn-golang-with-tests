package arrays_and_slices

func Sum(numbers [5]int) int {
	sum := 0
	for idx := 0; idx < len(numbers); idx++ {
		sum += numbers[idx]
	}
	return sum
}
