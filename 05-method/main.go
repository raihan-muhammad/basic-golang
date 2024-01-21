package main

import "fmt"

type biodata struct {
	name string
	age int8
	hobbies []string
}

type Biodata interface {
	getName() string
	getAge() int8
}

func initBio(bio Biodata){
	fmt.Print(bio.getName())
}

func (bio *biodata)getAge() int8 {
	return bio.age
}

func (bio *biodata)getName() string{
	return bio.name
}

func main(){
	person := biodata{
		name: "raihan",
		age: 23,
		hobbies: []string{ "mancing", "ngoding"},
	}
	initBio(&person)
	name := person.getName();
	fmt.Println(name)
}