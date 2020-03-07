package gofunc

// AddInt : Add 2 value toghether. Return currying function if only 1 params supplied
func AddInt(inputs ...int) int {
	result := 0
	for _, value := range inputs {
		result += value
	}

	return result
}

// AddString : all the string provided in the params together with a space
func AddString(inputs ...string) string {
	if len(inputs) == 1 {
		return inputs[0]
	}
	result := ""
	for index, value := range inputs {
		result += value
		if index != len(inputs)-1 { //Last params does not need to add space
			result += " "
		}
	}

	return result
}
