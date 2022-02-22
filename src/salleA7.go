package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)


// Salle permet d'afficher les choix disponibles dans la salleA7.
func Salle() {
	highlightTitle("Salle A7")
	fmt.Println()
	fmt.Print("Hind : ")
	SlowPrint("Bienvenue dans la salle A7! \n\n")
	SlowPrint(Colorize(Cyan, "L'endroit parfait pour se reposer un peu!\n\n"))
	fmt.Println("Jour", date, "\n")
	fmt.Print(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Dormir\n")
	fmt.Print(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Information du personnage\n")
	fmt.Print(Colorize(Bold, Colorize(LightMagenta, "3 - ")), "Contenu de l'inventaire\n")
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner au menu.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}


// Choixsalle permet de choisir une des options de la salleA7.
func (p *Personnage) Choixsalle() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Salle()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "0":
			p.Choixmenu()
		case "1":
			p.Dormir()
		case "2":
			p.ChoixInfo()
		case "3":
			p.DisplayInventory()
		default:
			p.Choixsalle()
		}
	}
}


// Dormir permet de restaurer ses PV et son manaa, dormir incrémente la date de 1.
func (p *Personnage) Dormir() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Rien de mieux qu'un bon sommeil pour se ressourcer! \n"))
	SlowPrint(Colorize(Cyan, "On explorera mieux demain..."))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Votre santé ainsi que vos manaas ont été restauré.")
	time.Sleep(2 * time.Second)
	date++
	p.manaa = p.manaamax
	p.pdvact = p.pdvmax
	p.Choixsalle()
}