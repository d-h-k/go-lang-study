package main

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// before refactoring : NOT use interface
/*
func printGreeting(englishBot englishBot) {
	println(englishBot.getGreeting())
}

func printGreeting(spanishBot spanishBot) {
	println(spanishBot.getGreeting())
}
*/
// before refactoring : use interface
func printGreeting(bot bot) {
	println(bot.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
