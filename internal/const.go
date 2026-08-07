package internal

type Command func(cmd []string)


var builtins = map[string]Command{}

func init() {
    builtins["echo"] = Echo
    builtins["type"] = Type
    builtins["pwd"] = Pwd
	builtins["exit"] = Exit
	builtins["cd"] = Cd
}



