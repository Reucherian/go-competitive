package main

import "fmt"

func main() {
	memo := map[int][]int{}
	inputArray := []int{2, 4}
	targetSum := 6
	fmt.Println(bestSum(inputArray, targetSum, memo))
}

func bestSum(inputArray []int, targetSum int, memo map[int][]int) []int {
	if _, found := memo[targetSum]; found {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	var bestResult []int
	for _, e := range inputArray {
		result := bestSum(inputArray, targetSum-e, memo)
		if result != nil {
			result = append([]int{e}, result...)
			if bestResult == nil || len(bestResult) > len(result) {
				bestResult = result
			}
		}
	}
	memo[targetSum] = bestResult
	return memo[targetSum]
}
