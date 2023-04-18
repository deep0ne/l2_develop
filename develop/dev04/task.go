package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(a, b int) bool {
		return bs[a] < bs[b]
	})
	return string(bs)
}

func FindAnagram(word string, anagrams map[string][]string) string {
	sortedWord := SortString(word)
	for key := range anagrams {
		sortedKey := SortString(key)
		if sortedWord == sortedKey {
			return key
		}
	}
	return ""
}

func PrettifyAnagrams(anagrams *map[string][]string) {
	for key, value := range *anagrams {
		if len(value) == 1 {
			delete(*anagrams, key)
			continue
		}
		sort.Strings(value)
		uniqueIndex := 0
		for i := 1; i < len(value); i++ {
			if value[i] != value[uniqueIndex] {
				uniqueIndex++
				value[uniqueIndex] = value[i]
			}
		}
		(*anagrams)[key] = value[:uniqueIndex+1]
	}
}

func FindAnagrams(words []string) map[string][]string {
	sort.Strings(words)
	anagrams := make(map[string][]string)
	for _, word := range words {
		word = strings.ToLower(word)
		anagram := FindAnagram(word, anagrams)
		if anagram == "" {
			anagrams[word] = append(anagrams[word], word)
		} else {
			anagrams[anagram] = append(anagrams[anagram], word)
		}
	}
	PrettifyAnagrams(&anagrams)
	return anagrams
}

func main() {
	words := []string{"пятак", "Пятак", "пятка", "тяпка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(FindAnagrams(words))
}
