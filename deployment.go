package main

// Deployment struct contain deployment specific informations
type Deployment struct {
	Overlays string
	Folder string
	Wait int
}

// Deps interface for Deployment struct
type Deps interface {
	setOverlays()
	setFolder()
	setWait()
}

func (d *Deployment) setOverlays() {

}