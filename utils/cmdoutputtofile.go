package utils

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func CmdOutputToFile(cmd *exec.Cmd, file string) {
	var cmdErr bytes.Buffer
	outfile, err := os.Create(file)
	if err != nil {
		log.Panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile
	cmd.Stderr = &cmdErr
	err = cmd.Run()
	if err != nil {
		log.Fatal(cmdErr.String())
	}
	cmd.Wait()
}
