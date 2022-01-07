package multiplicationtable

func MultiplicationTable(size int) [][]int {
	output := make([][]int, size)
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			output[i-1] = append(output[i-1], i*j)
		}
	}
	return output
}
