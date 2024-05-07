package main

import (
	"fmt"
	"tag-subfolder/structs"
)

func main() {
	fmt.Println(structs.Person{
		Name: "Bob",
		Age:  20,
	})
}
