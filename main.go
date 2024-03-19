package main

import (
	"flag"
	"fmt"
)

var (
	version bool
	name    bool
)

func main() {
	flag.BoolVar(&version, "v", false, "Display version")
	flag.BoolVar(&name, "n", false, "Display name")

	if name {
		user := getName("John")
		fmt.Println(user)
	}

	if version {
		fmt.Println(getVersion())
	}
}

func getName(name string) string {
	if len(name) == 0 {
		return "noname"
	} else {
		return fmt.Sprintf("Hi %s", name)
	}
}

func getVersion() string {
	return "1.0.0"
}
