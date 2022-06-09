package main

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	name  string
	words []string
	want  map[string][]string
}{
	{
		name:  "test-001",
		words: []string{"Пятак", "листок", "пятка", "столик", "столик", "тяпка", "слиток", "корова"},
		want:  map[string][]string{"пятак": {"пятка", "тяпка"}, "листок": {"слиток", "столик"}},
	},
	{
		name:  "test-002",
		words: []string{"замок", "МАЗОК", "кот", "ток", "кто", "аргумент", "кто", "Аргентум", "кто"},
		want:  map[string][]string{"замок": {"мазок"}, "аргумент": {"аргентум"}, "кот": {"кто", "ток"}},
	},
	{
		name:  "test-003",
		words: []string{"выход", "выдох", "костер", "корсет", "дверь", "сектор", "окрест"},
		want:  map[string][]string{"выход": {"выдох"}, "костер": {"корсет", "окрест", "сектор"}},
	},
}

func TestFindAnagrams(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := FindAnagrams(test.words)
			if !reflect.DeepEqual(res, test.want) {
				t.Fatal(fmt.Sprintf("Result: %s Want: %s", res, test.want))
			}
		})
	}
}
