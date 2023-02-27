package gomoduletest

import "fmt"

func Greeting(name string){
	fmt.Println("Hello",name)
}

func Greeting1(name string){
	fmt.Println("Hello",name)
}

type user struct{

	name string
	age int

}

type User1 struct{

	name string
	age int

}
