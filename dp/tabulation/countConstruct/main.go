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

	table := make([]int, len(target)+1)
	table[0] = 1

	for i := 0; i < len(table); i++ {
		if table[i] != 0 {
			for _, word := range wordBank {
				substringPrefix := string([]rune(target)[i:])
				if strings.HasPrefix(substringPrefix, word) && (i+len(word)) < len(table) {
					table[i+len(word)] += table[i]
				}
			}
		}
	}

	return table[len(target)]
}
