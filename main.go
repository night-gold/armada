package main

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
    "gopkg.in/src-d/go-git.v4/plumbing"
    "github.com/matthewrsj/copy"
)

func main() {
    filename := os.Args[1]
    var config Config
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }

    for _, repo := range config.Repo {
        var vers,gi,us,url,fo string
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
        url = gi + "/" + us + "/" + repo.Repository
        _, err := git.PlainClone("/tmp/"+ repo.Repository, false, &git.CloneOptions{
            URL: url,
            Progress: os.Stdout,
            ReferenceName: plumbing.ReferenceName("refs/heads/" + vers),
		    SingleBranch:  true,
        })
        if err != nil {
            panic(err)
        }

        erro := copy.All("/tmp/"+repo.Repository +"/base", fo + "/base")
        if erro != nil {
            panic(erro)
        }
    }
}