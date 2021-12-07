package main

import "fmt"

func Name(name string) string {
	if len(name) == 0 {
		return "noname"
	} else {
		return fmt.Sprintf("Hi %s", name)
	}

}
func main() {
	user := Name("John")
	fmt.Println(user)
}
