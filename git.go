package main

// Git struct contains all element refering to git
type Git struct {
	Repository string
	Git string
	Version string
	User string
	Private bool
}

// Repo interface is here to declare Git function
type Repo interface {
	setGit()
	setVersion()
	setUser()
}

func (g *Git) setGit() {
	if (g.Git == "") {
		g.Git = "https://github.com"
	}
}

func (g *Git) setVersion() {
	if (g.Version == "") {
		g.Version = "master"
	}
}

func (g *Git) setUser() {
	if (g.Git == "") {
		g.Git = "armada"
	}
}