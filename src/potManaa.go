package main

import (
	"fmt"
	"os"
	"strconv"
)


// ManaaPot permet de faire augmenter le manaa de 10.
func (p * Personnage) ManaaPot() {
	p.manaa += 10
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu gagnes 10 manaas!")
	fmt.Println("Tu as désormais ", Colorize(LightBlue, strconv.Itoa(p.manaa), " Manaa"))
}
