package internal

import "fmt"

func Execute(cmd []string) {
	if checkBuiltin(cmd) {
		builtins[cmd[0]](cmd)
		return
	}

	_, err := SearchBinary(cmd[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	External_exec(cmd)
}
