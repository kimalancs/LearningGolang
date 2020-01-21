package main

import "fmt"

// Peopler people
type Peopler interface {
	Speak(string) string
}

// Student student
type Student struct{}

// Speak speak
func (s *Student) Speak(think string) (talk string) {
	if think == "idiot" {
		talk = "You're so good at it"
	} else {
		talk = "hello"
	}
	return
}

func main() {
	var peo Peopler 
	fmt.Printf("%T\n", peo) // <nil>, 此时peo没有明确的类型
	peo = &Student{}
	fmt.Printf("%T\n", peo) // *main.Student, 此时peo有了明确的类型
	think := "bitch"
	fmt.Println(peo.Speak(think))

}
