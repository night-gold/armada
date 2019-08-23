package main

// Config is an array of Repo struct
type Config struct {
	Repo []Repo
}

// Repo represents a kustomize package app
type Repo struct {
	Repository string
	Git        string
	Version    string
	User       string
	Folder     string
	Overlays   string
	Private    bool
}
