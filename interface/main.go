package main

import "log"

type electricity interface {
	draw() string
}

type fan struct{}

func (fan) draw() string {
	return "Hello this is fan"
}

type kettle struct {
}

func (kettle) draw() string {
	return "Hello this is kettle"
}

func socket(e electricity) string {
	return e.draw()
}

func main() {
	var heyFan fan
	var heyKettle kettle
	log.Print(socket(heyFan))
	log.Print(socket(heyKettle))
	h := 3
	log.Print(&h)
}
