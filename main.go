package main

import (
	"fmt"
)

func main() {
	pokemon := Game{
		"Pokémon Rojo",
		"GAME FREAK",
		1990,
		64,
	}

	myGame := &pokemon

	info := (myGame).getInfo()
	fmt.Println(info)
}
