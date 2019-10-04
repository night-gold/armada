package main

import (
	"log"
)

// Git struct contains all element refering to git
type Git struct {
	Repository string
	Git        string
	Version    string
	User       string
	Private    bool
	Basepath   string
}

// Repo interface is here to declare Git function
type Repo interface {
	checkRepository()
	setGit()
	setVersion()
	setUser()
	setUrl()
	setBasePath()
}

func (g *Git) checkRepository() {
	if g.Repository == "" {
		log.Fatal("Git.Repository can't be empty!")
	}
}

func (g *Git) setGit() {
	if g.Git == "" {
		g.Git = "https://github.com"
	}
}

func (g *Git) setVersion() {
	if g.Version == "" {
		g.Version = "master"
	}
}

func (g *Git) setUser() {
	if g.Git == "" {
		g.Git = "armada"
	}
}

func (g *Git) setUrl() string {
	if g.Private {
		return g.Git + ":" + g.User + "/" + g.Repository
	}
	return g.Git + "/" + g.User + "/" + g.Repository
}

func (g *Git) setBasePath() {
	if g.Basepath == "" {
		g.Basepath = "."
	}
}
