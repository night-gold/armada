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
	var url string

	file := flag.String("f", "armada.yaml", "Armada package file to load")
	overlays := flag.String("o", "", "Default overlays for all deployment")
	apply := flag.Bool("a", false, "Auto apply of the kustomize configuration")
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
		if repo.Private {
			url = repo.Git + ":" + repo.User + "/" + repo.Repository
		} else {
			url = repo.Git + "/" + repo.User + "/" + repo.Repository
		}
		_, err := git.PlainClone(os.TempDir()+"/"+repo.Repository, false, &git.CloneOptions{
			URL:           url,
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

		if *apply {
			cmd := exec.Command("kubectl", "apply", "-k", "overlays/"+repo.Overlays)
			cmd.Dir = repo.Folder
			//output, err := cmd.CombinedOutput()
			err := cmd.Run()
			if err != nil {
				cleanFolder(repo.Repository, repo.Folder)
				log.Panic(err)
			}
		} else {
			cmd := exec.Command("kubectl", "kustomize", "overlays/"+repo.Overlays)
			cmd.Dir = repo.Folder

			if repo.Folder != repo.Repository && repo.Folder != "." {
				utils.CmdOutputToFile(cmd, repo.Folder+"-"+repo.Overlays+".yaml")
			} else {
				utils.CmdOutputToFile(cmd, repo.Repository+"-"+repo.Overlays+".yaml")
			}
		}

		//os.RemoveAll(os.TempDir() + "/" + repo.Repository)
		//os.RemoveAll(repo.Folder + "/base")
		cleanFolder(repo.Repository, repo.Folder)
	}
}

func cleanFolder(repo string, fold string) {
	os.RemoveAll(os.TempDir() + "/" + repo)
	os.RemoveAll(fold + "/base")
}
