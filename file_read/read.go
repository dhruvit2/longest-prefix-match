package fileread

import (
	"bufio"
	"errors"
	"os"
)

var (
	ErrInvalidFilePath = errors.New("Empty file path")
)

func ReadFileLineByLine(filePath string) ([]string, error) {
	var readFileData []string

	if filePath == "" {
		return readFileData, ErrInvalidFilePath
	}

	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		return readFileData, err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		readFileData = append(readFileData, fileScanner.Text())
	}

	return readFileData, nil
}
