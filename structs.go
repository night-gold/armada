package main

type Config struct {
    Repo []Repo
}

type Repo struct {
    Repository string
    Git string
    Version string
    User string
    Folder string
    Overlays string
}