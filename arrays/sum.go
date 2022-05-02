package main

func Sum(numberArray []int) (result int) {
	for _, number := range numberArray {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}
	return
}

func SumAllTails(numberToSum ...[]int) (result []int) {
	for _, numbers := range numberToSum {
		tail := []int{}
		if len(numbers) > 0 {
			tail = numbers[1:]
		}
		result = append(result, Sum(tail))
	}
	return
}

func main() {}