package main

import (
	"fmt"
	"os"
	"strconv"
)


// PoisonPot permet de diminuer le nombre de PV et affiche ceux restant à la fin.
func (p *Personnage) PoisonPot() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	for i := 0; i < 3; i++ {
		fmt.Println(Colorize(Red, "-10 PV"))
		p.pdvact -= 10
		if p.pdvact <= 0 {
			p.Dead()
			break
		}
	}
	fmt.Println("Il te reste désormais", Colorize(LightGreen, strconv.Itoa(p.pdvact), "PV"))
}


// ThrowPoisonPot permet de faire les effets de PoisonPot sur le mob ennemi.
func (c * Chose) ThrowPoisonPot() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	for i := 0; i < 3; i++ {
		fmt.Println(Colorize(Red, "-10 PV"))
		c.pdvact -= 10
		if c.pdvact <= 0 {
			if InHistoire {
				c.ChoseDeadHistoire()
			} else {
				c.ChoseDead()
			}
			break
		}
	}
	fmt.Println("Il reste ", Colorize(Gray, strconv.Itoa(c.pdvact), "PV"), " à", c.nom)
}
