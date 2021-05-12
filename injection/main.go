package main

import "log"

type Database interface {
	GetUser() string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "yoman"
}

type Greeter interface {
	Greet(name string)
}

type DefaultGreet struct{}

func (g DefaultGreet) Greet(name string) {
	log.Printf("Hello %s", name)
}

type Program struct {
	db Database
	gr Greeter
}

func (p Program) Exec() {
	user := p.db.GetUser()
	p.gr.Greet(user)
}

func main() {
	db := DefaultDatabase{}
	greeter := DefaultGreet{}
	p := Program{
		db: db,
		gr: greeter,
	}
	p.Exec()
}
