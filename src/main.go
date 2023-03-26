package main

import (
	"fmt"
	"game-of-life/print"
	"game-of-life/universe"
	"time"
)

func main() {
	unv := universe.NewUniverse(25)

	for true {
		gen := unv.NextGeneration()
		num := unv.NumberGeneration()
		cnt := unv.CountAlive()

		fmt.Print("\033[H\033[2J")
		fmt.Print(print.Printed(gen, num, cnt))

		time.Sleep(500 * time.Millisecond)
	}
}
