package main

// Package struct contains field refering to a specific package and one deployment
type Package struct {
	Name       string
	Deployment Deployment
}

// Pack set functions for package struct
type Pack interface {
	setDeployment(dep Deployment)
}

func (p *Package) setDeployment(overlays string) {
	p.Deployment.setOverlays(overlays)
}
