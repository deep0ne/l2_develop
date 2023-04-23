/*

Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
- Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
- Выходные данные: ссылка на мапу множеств анаграмм
- Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
- Массив должен быть отсортирован по возрастанию.
- Множества из одного элемента не должны попасть в результат.
- Все слова должны быть приведены к нижнему регистру.
- В результате каждое слово должно встречаться только один раз.

*/

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
		if len(value) == 1 { // Множества из одного элемента не должны попасть в результат.
			delete(*anagrams, key)
			continue
		}
		sort.Strings(value) // Массив должен быть отсортирован по возрастанию.
		uniqueIndex := 0
		for i := 1; i < len(value); i++ { // В результате каждое слово должно встречаться только один раз.
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
		word = strings.ToLower(word) // Все слова должны быть приведены к нижнему регистру.
		anagram := FindAnagram(word, anagrams)
		if anagram == "" {
			// Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого, слово из множества.
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
