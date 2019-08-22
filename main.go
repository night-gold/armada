package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/matthewrsj/copy"
	"github.com/night-gold/armada/utils"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/yaml.v2"
)

func main() {
	var config Config

	file := flag.String("f", "armada.yaml", "Armada package file to load")
	overlays := flag.String("o", "", "Default overlays for all deployment")
	flag.Parse()

	source, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Panic(err)
	}

	for _, repo := range config.Repo {
		if repo.Version == "" {
			repo.Version = "master"
		}
		if repo.Git == "" {
			repo.Git = "https://github.com"
		}
		if repo.User == "" {
			repo.User = "armada"
		}
		if repo.Folder == "" {
			repo.Folder = repo.Repository
		}
		if *overlays != "" && repo.Overlays == "" {
			repo.Overlays = *overlays
		} else if repo.Overlays == "" {
			repo.Overlays = "apply"
		}
		_, err := git.PlainClone(os.TempDir()+"/"+repo.Repository, false, &git.CloneOptions{
			URL:           repo.Git + "/" + repo.User + "/" + repo.Repository,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName("refs/heads/" + repo.Version),
			SingleBranch:  true,
		})
		if err != nil {
			log.Panic(err)
		}

		erro := copy.All(os.TempDir()+"/"+repo.Repository+"/base", repo.Folder+"/base")
		if erro != nil {
			log.Panic(erro)
		}

		cmd := exec.Command("kubectl", "kustomize", "overlays/"+repo.Overlays)
		cmd.Dir = repo.Folder

		if repo.Folder != repo.Repository && repo.Folder != "." {
			utils.CmdOutputToFile(cmd, repo.Folder+"-"+repo.Overlays+".yaml")
		} else {
			utils.CmdOutputToFile(cmd, repo.Repository+"-"+repo.Overlays+".yaml")
		}

		os.RemoveAll(os.TempDir() + "/" + repo.Repository)
		os.RemoveAll(repo.Folder + "/base")
	}
}
