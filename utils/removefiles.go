package utils

import (
	"log"
	"os"
)

func RemoveFiles(files []string) {
	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}
