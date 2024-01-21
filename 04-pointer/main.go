package main

import "fmt"

// pointer based by reference, not duplicate

type biodata struct {
	name string
	age int8
	isMerried bool
	hobbies []string
	school *School
}

type School struct {
	name string
	major []string
}
func main(){
	school := School{
		name: "smk muh1",
		major: []string { "RPL", "TKJ", "TEI"},
	}

	personA := biodata{
		name: "raihan",
		age: 23,
		isMerried: false,
		hobbies: []string { "mancing", "ngoding"},
		school: &school,
	}

	fmt.Println(personA.school.name);
}