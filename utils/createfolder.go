package utils

import (
	"os"
	"log"
)

// CreateFolder will create a folder with necessary parents
func CreateFolder(app string) {
	err := os.MkdirAll(app+"overlays/apply", 0755)
	if err != nil {
		log.Fatal("We couldn't create the folder")
	}
}