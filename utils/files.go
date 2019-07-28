package utils

import (
	"os"
	"os/exec"
)

// CmdOutputToFile write the cmd output to a file
func CmdOutputToFile(cmd *exec.Cmd, file string){
	outfile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile
	err = cmd.Start(); if err != nil {
		panic(err)
	}
	cmd.Wait()
}