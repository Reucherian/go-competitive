package main

import "fmt"

func main() {
	memo := map[int][]int{}
	inputArray := []int{4, 2}
	targetSum := 23
	fmt.Println(howSum(inputArray, targetSum, memo))
}

func howSum(inputArray []int, targetSum int, memo map[int][]int) []int {
	if _, found := memo[targetSum]; found {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return []int{-1}
	}
	for _, e := range inputArray {
		result := howSum(inputArray, targetSum-e, memo)
		if len(result) == 0 || len(result) >= 1 && result[0] != -1 {
			result = append([]int{e}, result...)
			memo[targetSum] = result
			return memo[targetSum]
		}
	}
	memo[targetSum] = []int{-1}
	return memo[targetSum]
}
