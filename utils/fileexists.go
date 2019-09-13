package utils

import (
	//	"errors"

	"fmt"
	"os"
)

// FileExists will test if file exists and return bool
func FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, fmt.Errorf("File does not exists")
	}

	if _, err := os.Stat(filePath); err == nil {
		//return true, errors.New("ok")
		return true, nil
	}

	return false, fmt.Errorf("We couldn't determined if file exists or not")
}
