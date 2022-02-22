package main

import (
	"bufio"
	"fmt"
	"os"
)


// shops affiche les choix possibles.
func (p *Personnage) shops() {
	highlightTitle("SHOPS")
	fmt.Println()
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Ben le marchand")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Anta le collectionneur")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 - ")), "Max le forgeron")
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner au menu.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}


// Choixshops permet de choisir une des options possibles liées aux shops.
func (p *Personnage) Choixshops() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	InFight = false
	counterRound = 0
	p.shops()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "0":
			p.Choixmenu()
		case "1":
			p.Choixmarchand()
		case "2":
			p.Collectionneur()
		case "3":
			p.Choixforgeron()
		default:
			p.Choixshops()
		}
	}
}
