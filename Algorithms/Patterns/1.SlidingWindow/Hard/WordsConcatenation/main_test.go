package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given a string and a list of words, find all the starting indices of substrings in the
 * given string that are a concatenation of all the given words exactly once without any
 * overlapping of words. It is given that all words are of the same length.
 *
 *	Example 1:
 *		Input: String="catfoxcat", Words=["cat", "fox"]
 *		Output: [0, 3]
 *		Explanation: The two substring containing both the words are "catfox" & "foxcat".
 *
 *	Example 2:
 *		Input: String="catcatfoxfox", Words=["cat", "fox"]
 *		Output: [3]
 *		Explanation: The only substring containing both the words is "catfox".
 *
 */

func findWordConcatenation(str string, words []string) []int {

	newIndices := []int{}
	wordsCount := len(words)
	wordLength := len(words[0])
	freqencyMap := make(map[string]int)

	// calculate frequqency of all words
	for _, word := range words {
		if _, ok := freqencyMap[word]; !ok {
			freqencyMap[word] = 0
		}
		freqencyMap[word] += 1
	}

	for i := 0; i <= (len(str) - wordsCount*wordLength); i++ {
		wordsSeen := make(map[string]int)

		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*wordLength
			// get the next word from the string
			word := str[nextWordIndex : nextWordIndex+wordLength]
			if _, ok := freqencyMap[word]; !ok {
				break // break if we don't need this word
			}

			// add the word to already seen map
			if _, ok := wordsSeen[word]; !ok {
				wordsSeen[word] = 0
			}
			wordsSeen[word] += 1

			// no need to process furthur if the word has frequency than required
			if wordsSeen[word] > freqencyMap[word] {
				break
			}

			// store the index we if have found all the words
			if j+1 == wordsCount {
				newIndices = append(newIndices, i)
			}
		}
	}

	return newIndices
}

func TestWordsConcatenation(t *testing.T) {

	type args struct {
		arg1 string
		arg2 []string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "String=catfoxcat Words=[cat, fox]",
			args: args{"catfoxcat", []string{"cat", "fox"}},
			want: []int{0, 3},
		},
		{
			name: "String=catcatfoxfox, Words=[cat, fox]",
			args: args{"catcatfoxfox", []string{"cat", "fox"}},
			want: []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWordConcatenation(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
