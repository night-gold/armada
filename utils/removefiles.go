package utils

import (
	"os"
)

func RemoveFiles(files []string) error {
	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			return err
		}
	}

	return nil
}
