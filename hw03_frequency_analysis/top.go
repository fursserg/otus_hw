package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	var uniqueWords []string
	wordsCount := make(map[string]int)

	for _, v := range strings.Fields(s) {
		if v == "-" {
			continue
		}

		v = strings.ToLower(strings.Trim(v, `:,.!"'`))

		if _, ok := wordsCount[v]; !ok {
			uniqueWords = append(uniqueWords, v)
		}

		wordsCount[v]++
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		if wordsCount[uniqueWords[i]] == wordsCount[uniqueWords[j]] {
			return uniqueWords[i] < uniqueWords[j]
		}

		return wordsCount[uniqueWords[i]] > wordsCount[uniqueWords[j]]
	})

	if len(uniqueWords) > 10 {
		return uniqueWords[:10]
	}

	return uniqueWords
}
