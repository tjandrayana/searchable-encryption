package transform

func Transform(input string) string {
	result := []byte{}

	// ini optional to make human can read this
	modifier := byte(100)

	for i := 1; i < len(input); i++ {
		result = append(result, input[i-1]-input[i]+modifier)
	}
	return string(result)
}
