package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello world!"
}

func (sb spanishBot) getGreeting() string {
	return "¡Hola Mundo!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	var eb englishBot
	var sb spanishBot
	printGreeting(eb)
	printGreeting(sb)
}
