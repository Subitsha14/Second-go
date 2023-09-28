package main

import (
	greetings "Greetings"
	"fmt"
	"log"
)

func main() {
	//caller function
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"subi", "sanju", "namjoon", "jin", "suga", "hobi", "tae", "jimin", "kookie"}
	greet, err := greetings.Hellos(names)
	//greet, err := greetings.Hello("subi")
	//greet, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greet)
}
