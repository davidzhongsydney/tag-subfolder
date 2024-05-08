package main

import (
	"fmt"

	"github.com/davidzhongsydney/tag-subfolder/structs"
)

func main() {
	fmt.Println(structs.Person{
		Name: "Bob",
		Age:  20,
	})
}
