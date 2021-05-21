package utils

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var templates []string
var overlay string

func Templating(vars []string, overlays string) []string {
	var removes []string
	overlay = overlays
	err := filepath.WalkDir(".", walkdir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range templates {
		t, err := template.ParseFiles(f)
		if err != nil {
			log.Fatal(err)
		}

		res, err := os.Create(strings.Replace(f, ".tmpl", "", -1))
		if err != nil {
			log.Fatal(err)
		}

		removes = append(removes, res.Name())

		config := make(map[string]string)
		for _, v := range vars {
			config[v] = os.Getenv(v)
		}

		err = t.Execute(res, config)
		if err != nil {
			log.Fatal(err)
		}

		res.Close()
	}

	return removes
}

func walkdir(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if strings.Contains(s, ".tmpl") && (strings.Contains(s, "base/") || strings.Contains(s, overlay)) {
		templates = append(templates, s)
	}
	return nil
}

func removeFiles(files []string) {
	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}
