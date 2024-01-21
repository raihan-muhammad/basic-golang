package main

import (
	Custom "basic-golang/helper"
	"fmt"
)


type biodata struct {
	name string
	age int8
	hobby []string
}


func main(){
	// example of variables
	name := "raihan"
	var age int8 = 23
	var isMerried bool = false
	fmt.Printf("age: %v isMerried: %v name: %s\n", age, isMerried, name)

	// example of slices
	hobby := []string { "mancing", "ngoding"}
	fmt.Println(hobby[1])

	// example of struct
	person := biodata {
		name: "raihan",
		age: 23,
		hobby: []string { "mancing", "ngoding"},
	}
	fmt.Println(person.hobby[0]);

	// example of map
	school := map[string]string{
		"name": "Smk mutuharjo",
		"classes": "100",
	}
	fmt.Println(school["classes"])
	
	fmt.Println(Custom.SayHi())
}
