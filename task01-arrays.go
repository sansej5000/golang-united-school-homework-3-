package homework

func average(input [15]float32) (result float32) {
	lenght := len(input)
	var sum float32
	for i := 0; i < lenght; i++ {
		sum += float32(input[i])
	}
	return sum / float32(lenght)
}
