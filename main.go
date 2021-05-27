package main

import (
	"bytes"
	"flag"
	"log"
	"os"
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
	remove := flag.Bool("r", false, "Remove the file generated using the templates")
	flag.Parse()

	packages.loadingYaml(*file)

	if len(packages.Namespaces) > 0 && *apply {
		createNamespace(packages.Namespaces)
	}

	for _, pack := range packages.Package {
		pack.setDeployment(*overlays)

		toDel, err := utils.Templating(pack.Deployment.Folder, pack.Variables, pack.Deployment.Overlays)
		if err != nil {
			log.Fatal(err)
		}

		if *apply {
			var cmdErr bytes.Buffer
			cmd := exec.Command("kubectl", "apply", "-k", "overlays/"+pack.Deployment.Overlays)
			cmd.Dir = pack.Deployment.Folder
			cmd.Stdout = os.Stdout
			cmd.Stderr = &cmdErr
			err := cmd.Run()
			if err != nil {
				log.Panic(cmdErr.String())
			}
		} else {
			cmd := exec.Command("kubectl", "kustomize", "overlays/"+pack.Deployment.Overlays)
			cmd.Dir = pack.Deployment.Folder

			if pack.Deployment.Folder != "" && pack.Deployment.Folder != "." && pack.Deployment.Folder != "../.." {
				utils.CmdOutputToFile(cmd, pack.Deployment.Folder+"-"+pack.Deployment.Overlays+".yaml")
			} else {
				utils.CmdOutputToFile(cmd, pack.Name+"-"+pack.Deployment.Overlays+".yaml")
			}
		}

		if *remove {
			utils.RemoveFiles(toDel)
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
