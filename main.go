package main

import (
	"fmt"
	readfile "github.com/dhruvit2/longest-prefix-match/file_read"
	search "github.com/dhruvit2/longest-prefix-match/prefix_search"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Invalid arguents, file name is required \n")
		return
	}

    if len(os.Args) < 3 {
		fmt.Printf("Invalid arguents, matching input string required \n")
		return
	}

	filePath := os.Args[1]
	searchStr := os.Args[2]

	fileLines, err := readfile.ReadFileLineByLine(filePath)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v", fileLines)

	result := search.ConcureAndSearch(fileLines, searchStr)

	if result.MaxMatchedCount == 0 {
		fmt.Printf("No matching string found \n")
        return
	}

	fmt.Printf("Longest matching string %v \n", result.LongestStr)
}
