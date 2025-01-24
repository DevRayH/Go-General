package main

import (
	parent "first_go/StructPractice/Struct1"
	child "first_go/StructPractice/Struct2"
	"fmt"
)

func main() {
	// Entry point

	child := child.Child{Parent: parent.Parent{Name : "Ray"}, Gender: "Male"}

	child.Speak()
	fmt.Println(child.Gender)
}