package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


// Fourrureloup permet d'acheter une Fourrure de loup chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) Fourrureloup() {
	if p.AddInventory(FdL) {
		p.money -= 4
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté de La Fourrure de loup! Retrouve la dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-4 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}


// Peaudetroll permet d'acheter une Peau de troll chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) Peaudetroll() {
	if p.AddInventory(PdT) {
		p.money -= 7
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté de La Peau de troll! Retrouve la dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-7 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}


// Cuirsanglier permet d'acheter du Cuir de sanglier chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) Cuirsanglier() {
	if p.AddInventory(CdS) {
		p.money -= 3
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté du Cuir de sanglier! Retrouve le dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-3 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}


// Plumecorbeau permet d'acheter une Plume de corbeau chez le marchand.
// L'objet sera ajouté à l'inventaire du personnage si le booléen de AddInventory est true (si l'inventaire a suffisamment de place).
// Le prix de l'objet sera déduit de l'argent du personnage.
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du marchand (cf. func Choixmarchand).
func (p *Personnage) Plumecorbeau() {
	if p.AddInventory(PdC) {
		p.money -= 2
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu as acheté une Plume de corbeau! Retrouve la dans ton inventaire.\n")
		fmt.Println(Colorize(Red, "-2 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		time.Sleep(3 * time.Second)
	}
	p.Choixmarchand()
}
