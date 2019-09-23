package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/matthewrsj/copy"
	"github.com/night-gold/armada/utils"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	var packages Packages

	file := flag.String("f", "armada.yaml", "Armada package file to load")
	overlays := flag.String("o", "", "Default overlays for all deployment")
	apply := flag.Bool("a", false, "Auto apply of the kustomize configuration")
	flag.Parse()

	packages.loadingYaml(*file)

	for _, pack := range packages.Package {
		pack.setGit()
		pack.setDeployment(*overlays)

		url := pack.Git.setUrl()

		ref := setRef(pack.Git.Version)
		_, err := git.PlainClone(os.TempDir()+"/"+pack.Git.Repository, false, &git.CloneOptions{
			URL:           url,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(ref + pack.Git.Version),
			SingleBranch:  true,
		})
		if err != nil {
			log.Panic(err)
		}

		erro := copy.All(os.TempDir()+"/"+pack.Git.Repository+"/base", pack.Deployment.Folder+"/base")
		if erro != nil {
			log.Panic(erro)
		}

		if *apply {
			cmd := exec.Command("kubectl", "apply", "-k", "overlays/"+pack.Deployment.Overlays)
			cmd.Dir = pack.Deployment.Folder
			//output, err := cmd.CombinedOutput()
			err := cmd.Run()
			if err != nil {
				cleanFolder(pack.Git.Repository, pack.Deployment.Folder)
				log.Panic(err)
			}
		} else {
			cmd := exec.Command("kubectl", "kustomize", "overlays/"+pack.Deployment.Overlays)
			cmd.Dir = pack.Deployment.Folder

			if pack.Deployment.Folder != pack.Git.Repository && pack.Deployment.Folder != "." {
				utils.CmdOutputToFile(cmd, pack.Deployment.Folder+"-"+pack.Deployment.Overlays+".yaml")
			} else {
				utils.CmdOutputToFile(cmd, pack.Git.Repository+"-"+pack.Deployment.Overlays+".yaml")
			}
		}

		cleanFolder(pack.Git.Repository, pack.Deployment.Folder)

		if pack.Deployment.Wait != 0 {
			time.Sleep(time.Duration(pack.Deployment.Wait) * time.Second)
		}
	}
}

func cleanFolder(repo string, fold string) {
	os.RemoveAll(os.TempDir() + "/" + repo)
	os.RemoveAll(fold + "/base")
}

func setRef(version string) string {
	ref := "refs/heads/"
	reg, err := regexp.MatchString(".*v[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.*", version)
	if err != nil {
		log.Panic(err)
	}
	if reg {
		ref = "refs/tags/"
		return ref
	}
	return ref
}
