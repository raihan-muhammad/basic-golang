package main

import (
	"fmt"
)

func main(){
	greeting, age := sayHi("raihan", 23)
	fmt.Println(greeting, age);
}

func sayHi(name string, age int8) (string, int8){
	return "hello world! my name is " + name, age
}