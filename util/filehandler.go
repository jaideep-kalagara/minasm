package util

import (
	"fmt"
	"os"
	"strings"
)

type FileHandler struct {
}

func (f FileHandler) SplitFile(filepath string) []string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fileformated := strings.ReplaceAll(string(file), "\r\n", "\n")
	return strings.Split(string(fileformated), "\n")
}
