package homework

func reverse(input []int64) (result []int64) {
	var results []int64
	len := len(input)
	for i := len; i > 0; i-- {
		results = append(results, int64(input[i-1]))
	}
	return results
}
