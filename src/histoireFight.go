package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


// mainfight fait commencer le combat selon l'initiative des personnages.
func mainfight(mob Chose) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	mob.DisplayInitChose()
	if mob.initiative > p1.initiative {
		fmt.Println(mob.nom, "commence le combat!")
		time.Sleep(2 * time.Second)
		p1.ChoseAttaque(mob)
	} else {
		p1.P1Attaque(mob)
	}
}


// P1Attaque génère l'attaque du personnage et enlève la possibilité de fuir (à l'inverse de PersoAttaque).
func (p *Personnage) P1Attaque(mob Chose) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	counterRound++
	ChoixfightHistoire()
	for {
		scannerFight := bufio.NewScanner(os.Stdin)
		scannerFight.Scan()
		optionsFight := scannerFight.Text()
		switch optionsFight {
		case "1":
			p.SkillUse(mob)
		case "2":
			p.invCombat(mob)
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.P1Attaque(mob)
		}
	}
}


// ChoixfightHistoire permet d'afficher les choix possibles dans le combat.
func ChoixfightHistoire() {
	fmt.Println("A toi de te fight!\n")
	fmt.Println("Tes manaas actuels : ", Colorize(LightBlue, strconv.Itoa(p1.manaa)), "\n")
	fmt.Println("Que décides-tu de faire?\n")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Compétence")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Inventaire\n")
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}
