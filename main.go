package main

import (
	"flag"
	"log"
	"os/exec"
	"time"

	"github.com/night-gold/armada/utils"
)

func main() {
	var packages Packages

	flag.Usage = flagUsage
	file := flag.String("f", "armada.yaml", "Armada package file to load")
	overlays := flag.String("o", "", "Default overlays for all deployment")
	apply := flag.Bool("a", false, "Auto apply of the kustomize configuration")
	flag.Parse()

	packages.loadingYaml(*file)

	if len(packages.Namespaces) > 0 {
		createNamespace(packages.Namespaces)
	}

	for _, pack := range packages.Package {
		pack.setDeployment(*overlays)

		if *apply {
			cmd := exec.Command("kubectl", "apply", "-k", "overlays/"+pack.Deployment.Overlays)
			cmd.Dir = pack.Deployment.Folder
			//output, err := cmd.CombinedOutput()
			err := cmd.Run()
			if err != nil {
				log.Panic(err)
			}
		} else {
			cmd := exec.Command("kubectl", "kustomize", "overlays/"+pack.Deployment.Overlays)
			cmd.Dir = pack.Deployment.Folder

			utils.CmdOutputToFile(cmd, pack.Deployment.Folder+"-"+pack.Deployment.Overlays+".yaml")
		}

		if pack.Deployment.Wait != 0 {
			time.Sleep(time.Duration(pack.Deployment.Wait) * time.Second)
		}
	}
}

func createNamespace(namespaces []string) {
	for _, namespace := range namespaces {
		cmd := exec.Command("kubectl", "create", "namespace", namespace)
		err := cmd.Run()
		if err != nil {
			log.Panic(err)
		}
	}
}
