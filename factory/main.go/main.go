package main

import (
	"fmt"
	"gocode/chapter10/factory/model"
)


func main() {

	student := model.GetStudent("tom",98.8)
	fmt.Printf("%p",&student)
	fmt.Println("name:",student.Name,"score:",student.GetScore())

	student1 := model.GetStudent("jdk",98.5)
	fmt.Printf("%p",&student1)	
}