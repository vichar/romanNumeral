package romanNumeral

func main() {
}

//Parse integer to roman numeral
func Parse(input int) string {
	output := ""
	if input > 0 && input <= 3 {
		for i := 1; i <= input; i++ {
			output = output + "I"
		}
	} else if input%4 == 0 {
		output = output + "IV"
	} else if (input < 10) && (input%5 == 0 || input%5 >= 1) {
		output = output + "V"
		repeat := input % 5
		if repeat < 4 {
			for i := 1; i <= repeat; i++ {
				output = output + "I"
			}
		} else {
			output = "IX"
		}

	} else if input%10 == 0 || input%10 >= 1 {
		output = output + "X"
		repeat := input % 10
		if repeat < 4 {
			for i := 1; i <= repeat; i++ {
				output = output + "I"
			}
		} else {

		}

	}

	return output
}
