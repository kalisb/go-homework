package main

func mapSum(function func(int) int, ints ...int) (result int) {
	for _, number := range ints {
		result += function(number)
	}
	return
}
