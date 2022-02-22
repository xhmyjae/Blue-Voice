package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


// AddSkills ajoute un sort à la liste des sorts du personnage.
func (p *Personnage) AddSkills(sorts string) {
	p.skill = append(p.skill, sorts)
	fmt.Println()
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as appris le skill", sorts, "\n")
	fmt.Println("Tes skills actuels : ", strings.Join(p.skill, ", "))
}


// BouledeFeu permet d'acheter le Sort : Boule de feu au marchand.
// Une fois acheté le prix sera déduit sur la money de personnage et le sort sera ajouté dans l'inventaire.
func (p *Personnage) BouledeFeu() {
	p.AddInventory("Sort : Boule de feu")
	p.money -= 25
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as acheté le Sort : Boule de feu! Retrouve le dans votre inventaire.\n")
	fmt.Println(Colorize(Red, "-25€"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.Choixmarchand()
}

