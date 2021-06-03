package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"text/template/parse"
)

var templates []string
var overlay string

func Templating(target string, configs map[string]string, overlays string) ([]string, error) {
	var removes []string
	config := configs
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

		node := listNodeFields(t.Tree.Root)
		var tNode []string
		for _, n := range node {
			var res string
			r := strings.Split(n, " ")
			for _, l := range r {
				if strings.Contains(l, ".") || strings.Contains(l, "{{.") {
					l = strings.Replace(l, "{{.", "", -1)
					l = strings.Replace(l, "}}", "", -1)
					l = strings.Trim(l, ".")
					res = strings.Trim(l, "$")
				}
			}
			if strings.Contains(n, "or ") {
				res = res + " or"
			}
			tNode = append(tNode, res)
		}

		fmt.Println(config)
		for _, n := range tNode {
			var value string
			var present bool
			if strings.Contains(n, " or") {
				n = strings.Replace(n, " or", "", -1)
				n = strings.Trim(n, " ")
				value, present = os.LookupEnv(n)
				if !present {
					fmt.Println("The variable " + n + " is not set but there is a default value in template " + t.Name() + "!")
				}
			} else {
				value, present = os.LookupEnv(n)
				if !present {
					return nil, errors.New("The variable " + n + " is not set but exists in the template " + t.Name() + "!")
				}
			}
			config[n] = value
		}

		fmt.Println(config)

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

func listNodeFields(node parse.Node) []string {
	var res []string
	if node.Type() == parse.NodeAction {
		res = append(res, node.String())
	}
	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			res = append(res, listNodeFields(n)...)
		}
	}
	return res
}
