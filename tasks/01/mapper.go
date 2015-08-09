package main

func Map(data []string, mapper func(string) string) []string {
	result := make([]string, 0)
	for _, value := range data {
		result = append(result, mapper(value))
	}
	return result
}

func Filter(data []string, predicate func(string) bool) []string {
	result := make([]string, 0)
	for _, value := range data {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func Reduce(data []string, combinator func(string, string) string) string {
	var result string
	for _, value := range data {
		result = combinator(result, value)
	}
	return result
}

func Any(data []string, predicate func(string) bool) bool {
	for _, value := range data {
		if predicate(value) {
			return true
		}
	}
	return false
}

func All(data []string, predicate func(string) bool) bool {
	for _, value := range data {
		if !predicate(value) {
			return false
		}
	}
	return true
}
