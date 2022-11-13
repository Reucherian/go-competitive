package main

import "fmt"

func main() {
	memo := map[int]bool{}
	inputArray := []int{2, 4}
	targetSum := 9
	fmt.Println(canSum(inputArray, targetSum, memo))
}

func canSum(inputArray []int, targetSum int, memo map[int]bool) bool {
	if _, found := memo[targetSum]; found {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, e := range inputArray {
		if canSum(inputArray, targetSum-e, memo) {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}
