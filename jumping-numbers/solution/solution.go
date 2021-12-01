package solution

import (
	"strconv"
)

func JumpingNumbers(n int, number int) interface{} {
	numbers := make([]int, 0)

	if n < 9 {
		for i := 0; i < 10; i++ {
			numbers = append(numbers, i)
		}
	}

	for i := 10; i < number; i++ {
		validation := validNumber(i, n)

		if validation != 0 {
			numbers = append(numbers, i)
		}
	}

	return numbers
}

func validNumber(value int, n int) int {
	sliceNumber := strconv.Itoa(value)
	var prev int32

	for index, number := range sliceNumber {
		if index == 0 {
			prev = number
			continue
		}

		if diff(number, prev) != int32(n) {
			return 0
		}

		prev = number
	}

	return value
}

func diff(a int32, b int32) int32 {
	if a < b {
		return b - a
	}
	return a - b
}
