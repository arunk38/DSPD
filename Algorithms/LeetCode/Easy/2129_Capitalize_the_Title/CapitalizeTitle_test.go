package capitalizeTitle_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * You are given a string title consisting of one or more words separated by a single space, where each word consists of English letters.
 * Capitalize the string by changing the capitalization of each word such that:
 * If the length of the word is 1 or 2 letters, change all letters to lowercase.
 * Otherwise, change the first letter to uppercase and the remaining letters to lowercase.
 *
 * Return the capitalized title.
 *
 * Example 1:
 *	Input: title = "capiTalIze tHe titLe"
 *	Output: "Capitalize The Title"
 *	Explanation:
 *		Since all the words have a length of at least 3, the first letter of each word is uppercase, and the remaining letters are lowercase.
 *
 */

func isLower(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

func capitalizeTitle(title string) string {

	titleToken := strings.Split(title, " ")
	for i, word := range titleToken {

		if len(word) < 3 {
			titleToken[i] = strings.ToLower(word)
			continue
		}
		if isLower(titleToken[i][0]) {
			titleToken[i] = string(titleToken[i][0]-('a'-'A')) + strings.ToLower(titleToken[i][1:])
		} else {
			titleToken[i] = string(titleToken[i][0]) + strings.ToLower(titleToken[i][1:])
		}
	}

	return strings.Join(titleToken, " ")
}

func TestLongestSubstringKDistinct(t *testing.T) {

	type args struct {
		arg1 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "title = [capiTalIze tHe titLe]",
			args: args{"capiTalIze tHe titLe"},
			want: "Capitalize The Title",
		},
		{
			name: "title = [First leTTeR of EACH Word]",
			args: args{"First leTTeR of EACH Word"},
			want: "First Letter of Each Word",
		},
		{
			name: "title = [i lOve leetcode]",
			args: args{"i lOve leetcode"},
			want: "i Love Leetcode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := capitalizeTitle(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
