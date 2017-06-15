package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	var max_len int

	for _, name := range names {
		if l := len(name); l > max_len {
			max_len = l
		}
	}

	output := make([][]string, max_len)

	for _, name := range names {
		output[len(name)-1] = append(output[len(name)-1], name)
	}

	fmt.Println("%v", output)

}
