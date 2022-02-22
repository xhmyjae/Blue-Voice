package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


// Forgeron affiche l'interface des choix possibles chez le forgeron : l'item à crafter et la possibilité de retourner aux shops.
func (p *Personnage) Forgeron() {
	highlightTitle("Maxime le forgeron")
	fmt.Print("\nMax : ")
	SlowPrint("T'as besoin d'un coup de main? Donne moi ce qui faut et je crafterai pour toi.\n\n")
	fmt.Println(Colorize(LightYellow, "Argent : ", strconv.Itoa(p.money), " Coins\n"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 - ")), CA, " : ", Colorize(LightYellow, "5 Coins"))
	fmt.Print(Colorize(Dim, "     Items nécessaires : Plume de corbeau (1), Cuir de sanglier (1).\n"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 - ")), TA, " : ", Colorize(LightYellow, "5 Coins"))
	fmt.Print(Colorize(Dim, "     Items nécessaires : Fourrure de loup (1), Peau de troll (2).\n"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 - ")), BA, " : ", Colorize(LightYellow, "5 Coins"))
	fmt.Print(Colorize(Dim, "     Items nécessaires : Fourrure de loup (1), Cuir de sanglier (1).\n"))
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner aux shops.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}


// Choixforgeron permet de sélectionner un des choix possibles et le faire s'exécuter.
// Exécute la function Forgeron avant de permettre à l'utilisateur de faire des choix.
// Si le personnage n'a pas l'argent requis, l'item ne pourra pas être crafté.
// La création des objets s'effectue via l'intermédiaire des functions dans le fichier affequipement.
func (p *Personnage) Choixforgeron() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.Forgeron()
	for {
		scannerForgeron := bufio.NewScanner(os.Stdin)
		scannerForgeron.Scan()
		options := scannerForgeron.Text()

		switch options {
		case "0":
			p.Choixshops()
		case "1":
			if p.money >= 5 {
				p.Chapeau()
			} else {
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
			}
		case "2":
			if p.money >= 5 {
				p.Tunique()
			} else {
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
			}
		case "3":
			if p.money >= 5 {
				p.Bottes()
			} else {
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
			}
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.Choixforgeron()
		}
	}
}
