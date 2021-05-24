package main

// Deployment struct contain deployment specific informations
type Deployment struct {
	Overlays string
	Folder   string
	Wait     int
}

// Deps interface for Deployment struct
type Deps interface {
	setOverlays()
	setFolder()
}

func (d *Deployment) setOverlays(overlays string) {
	if d.Overlays == "" && overlays == "" {
		d.Overlays = "apply"
	} else if overlays != "" {
		d.Overlays = overlays
	}
}

func (d *Deployment) setFolder() {
	if d.Folder == "" {
		d.Folder = "."
	} else if d.Overlays == "." {
		d.Folder = "../.."
	}
}
