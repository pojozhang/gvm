package machine

import "fmt"

type Command struct {
	JarFile string
	Args    []string
}

func Run(command *Command) {
	fmt.Println(command)
}
