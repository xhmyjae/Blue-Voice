package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


// DisplayInventory permet d'afficher les items présents dans l'inventaire en les numérotant.
// Si l'inventaire est vide un message dédié est affiché et le joueur est redirigé au menu.
// Sinon l'item demandé est utilisé via la function useItem.
func (p *Personnage) DisplayInventory() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	highlightTitle("Inventaire")
	fmt.Println()
	if len(p.inventaire) == 0 {
		fmt.Println(Colorize(Italic, "Ton inventaire est vide."))
		time.Sleep(2 * time.Second)
		p.Choixsalle()
	}

	counter := 1
	var arrItem []string
	for item, nbr := range p.inventaire {
		fmt.Println(Colorize(Bold, Colorize(LightMagenta, strconv.Itoa(counter), " - "))+item+" : ", nbr)
		arrItem = append(arrItem, item)
		counter++
	}
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner au menu.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	options := scanner.Text()
	number, err := strconv.Atoi(options)

	if err != nil {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInventory()
	}

	if number == 0 {
		p.Choixsalle()
	}

	if number-1 < len(arrItem) {
		item := arrItem[number-1]
		p.useItem(item)
		time.Sleep(2 * time.Second)
		p.DisplayInventory()
	} else {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInventory()
	}
}


// inInv place tous les éléments de l'inventaire dans un tableau et le revoie.
func (p *Personnage) inInv() []string {
	var items []string
	for item := range p.inventaire {
		items = append(items, item)
	}
	return items
}


// AugInventaire augmente la taille de l'inventaire après l'achat d'un item chez le marchant.
func (p *Personnage) AugInventaire() {
	p.limiteInv += 10
	p.money -= 30
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as acheté une Augmentation d'inventaire! Elle est directement implémentée à votre inventaire.\n")
	fmt.Println(Colorize(Red, "-30 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), "Coins"))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.Choixmarchand()
}
