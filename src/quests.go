package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var gain = []string{"		Gains : 10 Coins, 20 points d'exp.", "		Gains : 50 Coins, 45 points d'exp.", "		Gains : 200 Coins, 120 points d'exp.", "		Gains : 1000 Coins, 500 points d'exp.", "		Gains : Bonnet panda."}

var q1 = false
var q2 = false
var q3 = false
var q4 = false
var q5 = false


// Quests affiche les quêtes disponibles avec leurs gains.
func (p *Personnage) Quests() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Regarde! C'est les fameuses farming quests, les accomplir nous donnerais tellement d'expérience..."))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")

	highlightTitle("FARMING QUESTS")
	fmt.Println("\n")
	fmt.Println("  - Battre", c1.nom, "le 1er jour")
	if q1 == true {
		fmt.Println(Colorize(Gray, gain[0]))
	} else {
		fmt.Println(Colorize(Green, gain[0]))
	}
	fmt.Println("  - Atteindre le niveau 10")
	if q2 == true {
		fmt.Println(Colorize(Gray, gain[1]))
	} else {
		fmt.Println(Colorize(Green, gain[1]))
	}
	fmt.Println("  - Finir le mode histoire")
	if q3 == true {
		fmt.Println(Colorize(Gray, gain[2]))
	} else {
		fmt.Println(Colorize(Green, gain[2]))
	}
	fmt.Println("  - Battre", c2.nom)
	if q4 == true {
		fmt.Println(Colorize(Gray, gain[3]))
	} else {
		fmt.Println(Colorize(Green, gain[3]))
	}
	fmt.Println("  - Compléter et porter l'ensemble Nami (Haut pyjama Nami, Bas pyjama Nami)")
	if q5 == true {
		fmt.Println(Colorize(Gray, gain[4]))
	} else {
		fmt.Println(Colorize(Green, gain[4]))
	}
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner au menu.\n"))
	fmt.Print(Colorize(Bold, ">>> "))
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "0":
			p.Choixmenu()
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.Quests()
		}
	}
}


// quete1 si la condition est remplie, le bool en lien (q1) devient true.
func (p *Personnage) quete1() {
	if !q1 {
		if date == 1 {
			fmt.Println("Tu as validé la 1ère quête!")
			p.money += 10
			p.GainExp(20)
			q1 = true
		}
	}
}


// quete2 si la condition est remplie, le bool en lien (q2) devient true.
func (p *Personnage) quete2() {
	if !q2 {
		if p.niveau == 10 {
			fmt.Println("Tu as validé la 2ème quête!")
			p.money += 50
			p.GainExp(45)
			q2 = true
		}
	}
}


// quete3 si la condition est remplie, le bool en lien (q3) devient true.
func (p *Personnage) quete3() {
	if !q3 {
		if FinHistoire {
			fmt.Println("Tu as validé la 3ème quête!")
			p.money += 200
			p.GainExp(120)
			q3 = true
		}
	}
}


// quete4 si la condition est remplie, le bool en lien (q4) devient true.
func (p *Personnage) quete4() {
	if !q4 {
		fmt.Println("Tu as validé la 4ème quête!")
		p.money += 1000
		p.GainExp(500)
		q4 = true
	}
}


// quete5 si la condition est remplie, le bool en lien (q5) devient true.
func (p *Personnage) quete5() {
	if !q5 {
		if p.equipement.chestplate == HN && p.equipement.boots == BN {
			fmt.Println("Tu as validé la 5ème quête!")
			p.AddInventory(BPan)
			firstnami = true
			q5 = true
		}
	}
}
