package fizzbuzz

import "strconv"

func Play() []string {
	output := make([]string, 100)

	for i, _ := range output {
		value := i + 1
		dividesByThree := (value)%3 == 0
		dividesByFive := (value)%5 == 0

		if dividesByThree && dividesByFive {
			output[i] = "fizzbuzz"
		} else if dividesByThree {
			output[i] = "fizz"
		} else if dividesByFive {
			output[i] = "buzz"
		} else {
			output[i] = strconv.Itoa(value)
		}
	}

	return output
}
