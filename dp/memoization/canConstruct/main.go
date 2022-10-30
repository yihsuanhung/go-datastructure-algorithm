package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                                                    // true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                                     // false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))                                   // true
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // false
}

func canConstruct(target string, wordBank []string) bool {
	memo := make(map[string]bool)
	return mCanConsttruct(target, wordBank, memo)
}

func mCanConsttruct(target string, wordBank []string, memo map[string]bool) bool {
	// base case

	if m, ok := memo[target]; ok {
		return m
	}
	if target == "" {
		return true
	}

	// recursive
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			trimed := strings.TrimPrefix(target, word)
			if mCanConsttruct(trimed, wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}
