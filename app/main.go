package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/codecrafters-io/shell-starter-go/internal"
)


func main(){
	input := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Fprint(os.Stdout,"$ ")
		if !input.Scan(){
			break
		}
		cmd := input.Text()
		lexed := internal.Lexer(cmd)
		internal.Execute(lexed)
	}
}