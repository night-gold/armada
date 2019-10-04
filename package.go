package main

// Package struct contains field refering to a specific package and one deployment
type Package struct {
	Name       string
	Git        Git
	Deployment Deployment
}

// Pack set functions for package struct
type Pack interface {
	setGit(git Git)
	setDeployment(dep Deployment)
}

func (p *Package) setGit() {
	p.Git.checkRepository()
	p.Git.setGit()
	p.Git.setVersion()
	p.Git.setUser()
	p.Git.setBasePath()
}

func (p *Package) setDeployment(overlays string) {
	p.Deployment.setFolder(p.Git.Repository)
	p.Deployment.setOverlays(overlays)
}
