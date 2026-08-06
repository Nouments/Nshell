package internal

import (
	"os"
	"os/exec"
)


func External_exec(command []string) {
	if len(command)<2{
		cmd := exec.Command(command[0])
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
		return 
	}
	if 2<= len(command){
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
		return 
	}
}