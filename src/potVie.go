package main

import (
	"fmt"
	"os"
	"strconv"
)


// ViePot permet de rajouter 50 PV au Personnage si sa jauge de vie le permet et affiche les PV à la fin.
func (p *Personnage) ViePot() {
	if p.pdvact < p.pdvmax {
		p.pdvact += 50
	}
	if p.pdvmax <= p.pdvact {
		p.pdvact = p.pdvmax
	}
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as désormais ", Colorize(LightGreen, strconv.Itoa(p.pdvact), "/", strconv.Itoa(p.pdvmax) ,"PV"))
}
