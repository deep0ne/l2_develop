package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		words    []string
		anagrams map[string][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка"},
			anagrams: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			words: []string{"Просвещение", "Всепрощение", "Воспрещение"},
			anagrams: map[string][]string{
				"воспрещение": {"воспрещение", "всепрощение", "просвещение"},
			},
		},
		{
			words: []string{"ПЕСЧАНИК", "ПАСЕЧНИК", "ПЕСЧИНКА"},
			anagrams: map[string][]string{
				"пасечник": {"пасечник", "песчаник", "песчинка"},
			},
		},
		{
			words: []string{"вбгда", "дгвба", "абвгд", "гдбав", "бавдг", "гдавб", "гадвб", "вбгад", "гвадб", "бгдав"},
			anagrams: map[string][]string{
				"абвгд": {"абвгд", "бавдг", "бгдав", "вбгад", "вбгда", "гадвб", "гвадб", "гдавб", "гдбав", "дгвба"},
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := FindAnagrams(tc.words); !reflect.DeepEqual(output, tc.anagrams) {
			t.Errorf("QuickSort does not sort, %v not equal %v", output, tc.anagrams)
		}
	}

}
