package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func checkBuiltin(cmd []string) bool {
	_, exist := builtins[cmd[0]]
	return exist
}

func Echo(arg []string) {
	if len(arg) < 2 {
		fmt.Println("")
		return
	}
	fmt.Println(strings.Join(arg[1:], " "))
}

func Type(arg []string) {
	if len(arg) < 2 {
		fmt.Println("")
		return
	}
	if checkBuiltin(arg[1:]) {
		fmt.Printf("%s is a shell builtin\n", arg[1])
		return
	}
	PATH, err := SearchBinary(arg[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s is %s\n", arg[1], PATH)
}

func Pwd(arg []string) {
	cwd, err := filepath.Abs(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cwd)
}

func Exit(arg []string) {
	if len(arg) != 1 {
		fmt.Println("exit doesn't need argument")

	}
	os.Exit(0)
}

func Cd(arg []string) {
	if len(arg) != 2 {
		return
	}

	if filepath.IsAbs(arg[1]) {
		err:=os.Chdir(arg[1])
		if err!=nil{
			fmt.Printf("cd: %s: No such file or directory\n", arg[1] )
		}
		return
	}
	switch {
	case strings.HasPrefix(arg[1], "~"):
		homepath, err := os.UserHomeDir()
		if err != nil {
			return
		}
		if len((arg[1])) > 1 {
			homepath = filepath.Join(homepath, arg[1][2:])
		}
		err = os.Chdir(homepath)
		if err!=nil{
			fmt.Printf("cd: %s: No such file or directory\n", arg[1] )
		}
		return

	case strings.HasPrefix(arg[1], "./"):
		cwd, err := filepath.Abs(".")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir(filepath.Join(cwd, arg[1]))
		if err!=nil{
			fmt.Printf("cd: %s: No such file or directory\n", arg[1] )
		}

	default:
		path := filepath.Clean(arg[1])
		err := os.Chdir(path)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", arg[1] )
		}
	}
}
