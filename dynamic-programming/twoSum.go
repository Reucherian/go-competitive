package main

import "fmt"

func main() {
	var arr1 = []int{2, 3, 4, 5}
	arr := []int{2, 3, 5, 4}
	target := 7
	fmt.Println(twoSum(arr, target))
	fmt.Println(twoSum(arr1, target))
}

func twoSum(arr []int, target int) (int, int) {
	hash := map[int]int{}
	for _, ele := range arr {
		if _, found := hash[target-ele]; found {
			return target - ele, ele
		}
		hash[ele] = 1
	}
	return -1, -1
}
