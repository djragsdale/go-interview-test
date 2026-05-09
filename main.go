package main

import "fmt"

type App struct{}

func (a App) getGreeting() string {
	return "Hello World!"
}

func main() {
	app := App{}
	fmt.Println(app.getGreeting())
}
