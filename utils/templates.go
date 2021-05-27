package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var templates []string
var overlay string

func Templating(target string, vars []string, overlays string) ([]string, error) {
	var removes []string
	overlay = overlays
	err := filepath.WalkDir(target, walkdir)
	if err != nil {
		return nil, err
	}

	for _, f := range templates {
		res, err := os.Create(strings.Replace(f, ".tmpl", "", -1))
		if err != nil {
			return nil, err
		}

		t, err := template.New(filepath.Base(f)).Option("missingkey=error").ParseFiles(f)
		if err != nil {
			return nil, err
		}

		removes = append(removes, res.Name())

		var tmplVar []string
		validate := regexp.MustCompile(`^{{.*}}$`)
		for _, r := range t.Root.Nodes {
			if validate.MatchString(r.String()) {
				r1 := strings.Replace(r.String(), "{{", "", -1)
				r2 := strings.Replace(r1, "}}", "", -1)
				s := strings.Split(r2, " ")
				for _, res := range s {
					if strings.Contains(res, ".") {
						f := strings.Replace(res, ".", "", -1)
						tmplVar = append(tmplVar, f)
					}
				}
			}
		}

		config := make(map[string]string)
		for _, v := range vars {
			value, present := os.LookupEnv(v)
			if !present {
				return nil, errors.New("The variable " + v + " is not set!")
			}
			config[v] = value
		}

		for _, val := range tmplVar {
			if val, ok := config[val]; !ok {
				return nil, errors.New("The variable " + val + " is not set but exists in the template " + t.Name() + "!")
			}
		}

		err = t.Execute(res, config)
		if err != nil {
			return nil, err
		}

		res.Close()
	}

	return removes, nil
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
