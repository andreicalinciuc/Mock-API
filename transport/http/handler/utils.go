package handler

import (
	"os"
)

func checkIfFileExists(path string) bool {
	if _, err := os.Stat("data/" + path); err == nil {
		// path/to/whatever exists
		return true
	}

	return false
}
