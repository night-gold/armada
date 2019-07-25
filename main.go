package main

import (
    "fmt"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
	"io/ioutil"
    "os"
    "os/exec"
    "gopkg.in/src-d/go-git.v4/plumbing"
    "github.com/matthewrsj/copy"
)

var config Config

func init() {
    filename := os.Args[1]
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }
}

func main() {
    for _, repo := range config.Repo {
        var vers,gi,us,url,fo,ov string
        if repo.Version != ""{
           vers = repo.Version
        }else{
            vers = "master"
        }
        if repo.Git != ""{
            gi = repo.Git
        }else{
            gi = "https://github.com"
        }
        if repo.User != ""{
            us = repo.User
        }else{
            us = "armada"
        }
        if repo.Folder != ""{
            fo = repo.Folder
        }else{
            fo = repo.Repository
        }
        if repo.Overlays != ""{
            ov = repo.Overlays
        }else{
            ov = "apply"
        }
        url = gi + "/" + us + "/" + repo.Repository
        _, err := git.PlainClone(os.TempDir()+"/"+ repo.Repository, false, &git.CloneOptions{
            URL: url,
            Progress: os.Stdout,
            ReferenceName: plumbing.ReferenceName("refs/heads/" + vers),
		    SingleBranch:  true,
        })
        if err != nil {
            panic(err)
        }

        erro := copy.All(os.TempDir()+"/"+repo.Repository +"/base", fo + "/base")
        if erro != nil {
            panic(erro)
        }

        cmd, errors := exec.Command("kustomize build overlays/"+ ov +" > "+ repo.Repository +".yaml").Output()
        if errors != nil {
            panic(cmd)
        }
        fmt.Println(cmd)

        os.RemoveAll(os.TempDir()+"/"+repo.Repository)
        os.RemoveAll(fo+"/base")
    }
}