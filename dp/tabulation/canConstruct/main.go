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

	table := make([]bool, len(target)+1)
	table[0] = true

	for i := 0; i < len(table); i++ {
		if table[i] {
			for _, word := range wordBank {
				substringPrefix := string([]rune(target)[i:])
				if strings.HasPrefix(substringPrefix, word) && (i+len(word)) < len(table) {
					table[i+len(word)] = true
				}
			}
		}
	}

	return table[len(target)]
}
