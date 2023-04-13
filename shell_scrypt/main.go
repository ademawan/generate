package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "./cd_test.sh", "~/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	fmt.Println(cmd.Path)
	// output := string(cmd)
	// fmt.Println(output)
	cmd.Run()
	command := []string{
		"./cd_test.sh",
		"arg1=~/",
	}

	Execute("./cd_test.sh", command)

}

func Execute(script string, command []string) (bool, error) {

	cmd := &exec.Cmd{
		Path:   script,
		Args:   command,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err := cmd.Start()
	if err != nil {
		return false, err
	}

	err = cmd.Wait()
	if err != nil {
		return false, err
	}
	fmt.Println(cmd.Output())
	return true, nil
}
