package main

import "fmt"

func main() {

	people := map[string]string{
		"Salman": "Swidish",
		"Marvan": "UAE",
		"Zobair": "Iranian",
	}

	printMap(people)

}

func printMap(c map[string]string) {
	for person, nanionality := range c {
		fmt.Println(person+" : ", nanionality)
	}
}
