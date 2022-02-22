package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


// PotionDeVie permet d'acheter une Potion de vie chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) PotionDeVie() {
	if p.AddInventory(PtV) {
		p.money -= 3
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté une Potion de vie! Retrouve la dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-3 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}


// PotionDePoison permet d'acheter une Potion de poison chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) PotionDePoison() {
	if p.AddInventory(PtP) {
		p.money -= 6
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté une Potion de poison! Retrouve la dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-6 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}


// PotionDeManaa permet d'acheter une Potion de Mana chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) PotionDeManaa() {
	if p.AddInventory(PtM) {
		p.money -= 10
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté une Potion de manaa! Retrouve la dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-10 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}
