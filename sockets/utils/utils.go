package utils

import (
	"fmt"
	"os"
)

func SaveStringFile(path string, data []byte) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error while creating file:", err)
	}
	defer file.Close()
	_, err = file.WriteString(string(data))
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
