package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                                                    // 2
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))                                         // 1
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                                     // 0
	fmt.Println(allConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // 0
}

func allConstruct(target string, wordBank []string) [][]string {
	memo := make(map[string][][]string)
	return mAllConstruct(target, wordBank, memo)
}

func mAllConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {

	if m, ok := memo[target]; ok {
		return m
	}
	if target == "" {
		return [][]string{{}}
	}

	result := [][]string{}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			trimed := strings.TrimPrefix(target, word)
			ways := mAllConstruct(trimed, wordBank, memo)
			for i, way := range ways {
				ways[i] = append([]string{word}, way...)
			}
			result = append(result, ways...)
		}
	}

	memo[target] = result
	return result
}
