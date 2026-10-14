package lesson2

func FizzBuzz(n int) string {
	switch {
	case n <= 0:
		return "invalid"
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	}

	return ""
}

func Sum(values []int64) int64 {
	if values == nil {
		return 0
	}

	if len(values) == 0 {
		return 0
	}

	var sum int64

	for _, value := range values {
		sum += value
	}

	return sum
}

func Min(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}

	min := values[0]

	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min
}

func Max(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}

	max := values[0]

	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max
}

func Average(values []int64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum int64

	var length int = len(values)

	for i := range length {
		sum += values[i]
	}

	var average float64 = float64(sum) / float64(length)

	return average
}

func Category(amount int64) string {
	switch {
	case amount < 0:
		return "invalid"
	case amount <= 9999:
		return "small"
	case amount <= 49999:
		return "regular"
	}

	return "large"
}

func RuneCount(text string) int {
	var cnt int = 0

	for range text {
		cnt++
	}

	return cnt
}
