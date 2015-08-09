package main

import (
	"sort"
	"strings"
)

func getMostCommonWords(text string, wordCountLimit int) []string {

	result := getResult(strings.Split(strings.ToLower(text), " "), wordCountLimit)
	sort.Strings(result)
	return result
}

func getResult(text []string, count int) []string {
	result := make([]string, 0)
	mapCount := getCount(text)
	for key, value := range mapCount {
		if value >= count {
			result = append(result, key)
		}
	}
	return result
}

func getCount(text []string) map[string]int {
	mapCount := make(map[string]int)
	for _, value := range text {
		if _, ok := mapCount[value]; ok {
			mapCount[value]++
		} else {
			mapCount[value] = 1
		}
	}
	return mapCount
}
