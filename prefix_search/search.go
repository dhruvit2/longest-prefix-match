package search

import (
	model "github.com/dhruvit2/longest-prefix-match/store"
)

func matchPrefix(inputStr string, match string) int {
	inputStrLen := len(inputStr)
	matchStrLen := len(match)
	count := 0
	for i := 0; i < inputStrLen && i < matchStrLen; i++ {
		if inputStr[i] == match[i] {
			count++
		} else {
			break
		}
	}

	return count
}

func ConcureAndSearch(arrStrings []string, match string) model.Result {
	if len(arrStrings) <= 1 {
		count := matchPrefix(arrStrings[0], match)
		if count == 0 {
			return model.Result{LongestStr: "", MaxMatchedCount: count}
		}
		return model.Result{LongestStr: arrStrings[0], MaxMatchedCount: count}
	}
	helfLength := len(arrStrings) / 2
	leftResult := ConcureAndSearch(arrStrings[:helfLength], match)
	rightResult := ConcureAndSearch(arrStrings[helfLength:], match)
	if leftResult.MaxMatchedCount > rightResult.MaxMatchedCount {
		return leftResult
	} else if rightResult.MaxMatchedCount > leftResult.MaxMatchedCount {
		return rightResult
	} else if rightResult.MaxMatchedCount == leftResult.MaxMatchedCount {
		if len(rightResult.LongestStr) > len(leftResult.LongestStr) {
			return rightResult
		} else {
			return leftResult
		}
	}
	return model.Result{}
}
