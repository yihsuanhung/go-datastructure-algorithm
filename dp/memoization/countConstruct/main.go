package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                                                    // 2
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                                                    // 1
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                                     // 0
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))                                   // 4
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // 0
}

func countConstruct(target string, wordBank []string) int {
	memo := make(map[string]int)
	return mCountConstruct(target, wordBank, memo)
}

func mCountConstruct(target string, wordBank []string, memo map[string]int) int {
	// base case
	if m, ok := memo[target]; ok {
		return m
	}
	if target == "" {
		return 1
	}

	result := 0

	// recursive
	for _, word := range wordBank {

		if strings.HasPrefix(target, word) {
			trimed := strings.TrimPrefix(target, word)
			result += mCountConstruct(trimed, wordBank, memo)
		}
	}

	memo[target] = result
	return result
}
