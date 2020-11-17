package utils

import (
	"strings"
)

// GetTargetFolder will return the last folder in the Path given
func GetTargetFolder(filePath string) (string) {
	if strings.Contains(filePath, "/") {
		folder := strings.Split(filePath, "/")
		folderLen := len(folder)-1
		return folder[folderLen]
	}else{
		return filePath
	}
}
