package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


// Marchand affiche l'interface des choix possibles chez le marchand : items disponibles à l'achat et la possibilité de retourner aux shops.
func (p *Personnage) Marchand() {
	highlightTitle("Ben le Marchand")
	fmt.Println()
	fmt.Print("Ben : ")
	SlowPrint("Piste la marchandise, c'est de la qualité. Ca t'intéresse? \n\n")
	fmt.Println(Colorize(LightYellow, "Argent : ", strconv.Itoa(p.money), " Coins\n"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Potion de vie : ", Colorize(LightYellow, "3 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Potion de poison : ", Colorize(LightYellow, "6 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 - ")), "Potion de manaa : ", Colorize(LightYellow, "10 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "4 - ")), "Sort : Boule de feu : ", Colorize(LightYellow, "25 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "5 - ")), "Fourrure de loup : ", Colorize(LightYellow, "4 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "6 - ")), "Peau de troll : ", Colorize(LightYellow, "7 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "7 - ")), "Cuir de sanglier : ", Colorize(LightYellow, "3 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "8 - ")), "Plume de corbeau : ", Colorize(LightYellow, "2 Coins"))
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "9 - ")), "Augmentation d'inventaire : ", Colorize(LightYellow, "30 Coins"))
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner aux shops.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}


// Choixmarchand permet de sélectionner un des choix possible et le faire s'exécuter.
// Si le personnage n'a pas l'argent requis, l'item ne pourra pas être acheté.
// L'achat d'objet s'effectue via les functions propres à chaque item.
func (p *Personnage) Choixmarchand() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.Marchand()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "0":
			p.Choixshops()
		case "1":
			if p.money >= 3 {
				p.PotionDeVie()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "2":
			if p.money >= 6 {
				p.PotionDePoison()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "3":
			if p.money >= 10 {
				p.PotionDeManaa()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "4":
			if inList("Boule de feu", p.skill) || inList("Sort : Boule de feu", p.inInv()) {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				fmt.Println("Tu connais déjà ce sort.")
				time.Sleep(1 * time.Second)
			} else {
				if p.money >= 25 {
					p.BouledeFeu()
				} else {
					os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
					SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
					time.Sleep(1 * time.Second)
				}
			}
		case "5":
			if p.money >= 4 {
				p.Fourrureloup()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "6":
			if p.money >= 7 {
				p.Peaudetroll()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "7":
			if p.money >= 3 {
				p.Cuirsanglier()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "8":
			if p.money >= 2 {
				p.Plumecorbeau()
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
				time.Sleep(1 * time.Second)
			}
		case "9":
			if p.limiteInv < 40 {
				if p.money >= 30 {
					p.AugInventaire()
				} else {
					os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
					SlowPrint(Colorize(Cyan, "Frero t'es vraiment pauvre"))
					time.Sleep(1 * time.Second)
				}
			} else {
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				fmt.Println("Tu as atteint la limite d'inventaire!")
				time.Sleep(1 * time.Second)
			}
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
		}
		p.Choixmarchand()
	}
}
