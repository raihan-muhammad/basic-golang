package main

import "fmt"

type biodata struct{
	name string
	age int8
	isMerried bool
	hobbies []string
}

func main(){
	person := biodata{
		name: "raihan",
		age: 23,
		isMerried: false,
		hobbies: []string{ "mancing", "watching football"},
	}

	if person.isMerried  {
		fmt.Println("I already merried")
	} else {
		fmt.Println("I am not merried")
	}

	if person.age > 23 {
		fmt.Println("you are old!")
	} else {
		fmt.Println("you are young!")
	}

	for _, hobby := range person.hobbies {
		fmt.Println(hobby)
	}
	
	for i:=0; i < len(person.hobbies); i++ {
		fmt.Println(person.hobbies[i]);
	}
}