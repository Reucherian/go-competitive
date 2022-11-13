package main

import "fmt"

func main() {
	fmt.Println(canConstruct([]string{"ab", "bcd", "cef"}, "abcdef", map[string]bool{}))
	fmt.Println(canConstruct([]string{"bo", "rt", "ate", "t", "ska", "sk", "boar"}, "skateboard", map[string]bool{}))
	fmt.Println(canConstruct([]string{"a", "p", "ent", "enter", "ot", "o", "t"}, "enterapotentpot", map[string]bool{}))
	fmt.Println(canConstruct([]string{"e", "ee", "eee", "eeee", "eeeeee", "eeeeeeee"}, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", map[string]bool{}))

}

func canConstruct(stringList []string, finalString string, memo map[string]bool) bool {
	if _, found := memo[finalString]; found {
		return memo[finalString]
	}
	if finalString == "" {
		return true
	}
	for _, e := range stringList {
		if checkPreWord(e, finalString) {
			if canConstruct(stringList, finalString[len(e):], memo) {
				memo[finalString] = true
				return true
			}
		}

	}
	memo[finalString] = false
	return false
}

func checkPreWord(word string, targetWord string) bool {
	for i, e := range word {
		if byte(e) != targetWord[i] {
			return false
		}
	}
	return true
}
