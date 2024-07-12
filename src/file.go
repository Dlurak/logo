package src

import (
	"os"
	"strings"
)

func writeToFile(filepath string, line string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	if strings.HasSuffix(line, "\n") {
		file.WriteString(line)
	} else {
		file.WriteString(line + "\n")
	}
}
