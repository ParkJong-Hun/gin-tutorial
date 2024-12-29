package utility

import (
	"fmt"
	"os"
)

func MkDirIfNeeded(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("mkDir error:", err)
			return false
		}
		return true
	}
	return false
}
