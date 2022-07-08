package main

import "fmt"

func main() {
	user := Name("John")
	fmt.Println(user)
}

func Name(name string) string {
	if len(name) == 0 {
		return "noname"
	} else {
		return fmt.Sprintf("Hi %s", name)
	}

}
