package main

func Map(data []string, mapper func(string) string) (result []string) {
	result = make([]string, len(data))
	for i, value := range data {
		result[i] = mapper(value)
	}
	return
}

func Filter(data []string, predicate func(string) bool) (result []string) {
	result = make([]string, 0, len(data))
	for _, value := range data {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func Reduce(data []string, combinator func(string, string) string) (result string) {
	if len(data) == 1 {
		return data[0]
	}
	for _, value := range data {
		result = combinator(result, value)
	}
	return
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
