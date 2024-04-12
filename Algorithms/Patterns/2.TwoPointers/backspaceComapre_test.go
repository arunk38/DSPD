package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given two strings containing backspaces (identified by the character ‘#’),
 * check if the two strings are equal.
 */

func getNextValidCharIndex(str string, index int) int {
	backSpaceCount := 0

	for index >= 0 {
		if str[index] == '#' { // found a backspace char
			backSpaceCount++
		} else if backSpaceCount > 0 { // a non-backspace char
			backSpaceCount--
		} else {
			break
		}

		index-- // skip a backspace for valid char
	}

	return index
}

func compareBackspaceStrings(str1 string, str2 string) bool {
	index1 := len(str1) - 1
	index2 := len(str2) - 1

	for index1 >= 0 || index2 >= 0 {
		i1 := getNextValidCharIndex(str1, index1)
		i2 := getNextValidCharIndex(str2, index2)

		if i1 < 0 && i2 < 0 { // reached end of both strings
			return true
		}

		if i1 < 0 || i2 < 0 { // reached end of only one string
			return true
		}

		if str1[i1] != str2[i2] { // check if the chars are euqal
			return false
		}

		index1 = i1 - 1
		index2 = i2 - 1
	}
	return true
}

func TestCompareBackspaceStrings(t *testing.T) {

	type args struct {
		arg1 string
		arg2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "str1=xy#z, str2=xzz#",
			args: args{"xy#z", "xzz#"},
			want: true,
		},
		{
			name: "str1=xy#z, str2=xyz#",
			args: args{"xy#z", "xyz#"},
			want: false,
		},
		{
			name: "str1=xp#, str2=xyz##",
			args: args{"xp#", "xyz##"},
			want: true,
		},
		{
			name: "str1=xywrrmp, str2=xywrrmu#p",
			args: args{"xywrrmp", "xywrrmu#p"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareBackspaceStrings(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
