package main

import (
	"io/ioutil"
	"log"

	"github.com/night-gold/armada/utils"
	"gopkg.in/yaml.v2"
)

// Packages is an array of Packages struct
type Packages struct {
	Package []Package
}

// Pac set interface for packages structs
type Pac interface {
	loadingYaml()
}

func (p *Packages) loadingYaml(file string) {
	a, err := utils.FileExists(file)
	if a == false {
		log.Fatal(err)
	}
	source, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &p)
	if err != nil {
		log.Panic(err)
	}
}
