package main

// Packages is an array of Packages struct
type Packages struct {
	Package []Package
}

// Package struct contains field refering to a specific package and one deployment
type Package struct {
	Git Git
	Deployment Deployment
}

// Pack set functions for package struct
type Pack interface {
	setGit(git Git)
	setDeployment(dep Deployment)
}

func (p *Package) setGit() {
	p.Git.setGit()
	p.Git.setVersion()
	p.Git.setUser()
}

func (p *Package) setDeployment(overlays string) {
	p.Deployment.setFolder(p.Git.Repository)
	p.Deployment.setOverlays(overlays)
}