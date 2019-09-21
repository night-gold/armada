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
}

func (d *Deployment) setOverlays() {
	if(d.Overlays == ""){
		d.Overlays = "apply"
	}
}

func (d *Deployment) setFolder(repo string) {
	if(d.Folder == ""){
		d.Folder = repo
	}
}