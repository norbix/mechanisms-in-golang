package main

import (
	"log"
)

//NOTE: <<interface>>
type Database interface {
	GetUser() string
	GetAllUsers() []string
}

//NOTE: <component>
type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "norbix"
}

//NOTE: <component>
type FamousDatabase struct{}

func (db FamousDatabase) GetUser() string {
	return "Fresh Prince of Bel Air"
}

func (db FamousDatabase) GetAllUsers() []string {
	return []string{}
}

//NOTE: <<interface>>
type Greeter interface {
	Greet(userName string)
}

//NOTE: <component>
type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s!! Nice 2 meet u", userName)
}

//NOTE: <<component>>
type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

// NOTE: entrypoint
func main() {
	//db := DefaultDatabase{}
	db := FamousDatabase{}
	greeter := NiceGreeter{}
	p := Program{
		Db:      db,
		Greeter: greeter,
	}

	p.Execute()
}
