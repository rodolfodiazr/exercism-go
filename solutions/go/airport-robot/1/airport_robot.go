package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(visitorsName string) string
}

func SayHello(visitorsName string, greeter Greeter) string {
    return fmt.Sprintf("I can speak %v: %v", greeter.LanguageName(), greeter.Greet(visitorsName))
}

type Italian struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(visitorsName string) string {
    return "Ciao " + visitorsName + "!"
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(visitorsName string) string {
    return "Olá " + visitorsName + "!"
}