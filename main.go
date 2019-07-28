package main

import (
    "flag"
    git "gopkg.in/src-d/go-git.v4"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "os/exec"
    "gopkg.in/src-d/go-git.v4/plumbing"
    "github.com/matthewrsj/copy"
    "github.com/night-gold/armada/utils"
)

func main() {
    var config Config

    file := flag.String("f", "armada.yaml", "Armada package file to load")
    help := flag.Bool("h", false, "Print the help")
    flag.Parse()
    
    if *help {
		flag.PrintDefaults()
	} else {
        source, err := ioutil.ReadFile(*file)
        if err != nil {
            panic(err)
        }
        err = yaml.Unmarshal(source, &config)
        if err != nil {
            panic(err)
        }

        for _, repo := range config.Repo {
            /* var vers,gi,us,url,fo,ov string */
            if repo.Version == ""{
                repo.Version = "master"
            }
            if repo.Git == ""{
                repo.Git = "https://github.com"
            }
            if repo.User == ""{
                repo.User = "armada"
            }
            if repo.Folder != ""{
                repo.Folder = repo.Repository
            }
            if repo.Overlays != ""{
                repo.Overlays = "apply"
            }
            /* url = gi + "/" + us + "/" + repo.Repository */
            _, err := git.PlainClone(os.TempDir()+"/"+ repo.Repository, false, &git.CloneOptions{
                URL: repo.Git + "/" + repo.User + "/" + repo.Repository,
                Progress: os.Stdout,
                ReferenceName: plumbing.ReferenceName("refs/heads/" + repo.Version),
                SingleBranch:  true,
            })
            if err != nil {
                panic(err)
            }

            erro := copy.All(os.TempDir()+"/"+repo.Repository +"/base", repo.Folder + "/base")
            if erro != nil {
                panic(erro)
            }

            cmd := exec.Command("kubectl","kustomize", "overlays/"+ repo.Overlays)
            utils.CmdOutputToFile(cmd, repo.Repository + ".yaml")

            os.RemoveAll(os.TempDir()+"/"+repo.Repository)
            os.RemoveAll(repo.Folder+"/base")
        }
    }
}