package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct{}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
func (i Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}
