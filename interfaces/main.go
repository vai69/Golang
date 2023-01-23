package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreenting(eb)
	printGreenting(sb)
}

func printGreenting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hii there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}