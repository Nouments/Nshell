package internal

import(
	"fmt"
	"strings"
)

type Command struct {
	Name string
}



builtins := []Command{
		{Name: "echo"},
		{Name: "exit"},
		{Name: "type"},
		{Name: "pwd"},
}

func checkBuiltin(cmd string) bool{
	for _,builtin :=  range builtins {
		if cmd == builtin {
			return true
		}
	}
	return false
}


func Echo(arg []string){
	fmt.Println(strings.Join(arg," "))
}

func Type(string)