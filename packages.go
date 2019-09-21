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