package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


// Collectionneur permet de dynamiser l'interface de vente.
// Si le personnage n'a pas d'objets dans son inventaire il sera revoyé automatiquement à Choixshops
// Sinon les items à vendre seront affichés.
func (p *Personnage) Collectionneur() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	highlightTitle("Anta le rat")
	fmt.Print("\nAnta : ")
	SlowPrint("Si t'as quelques choses à vendre t'es au bon endroit! \n\n")
	if len(p.inventaire) == 0 {
		for i := 0; i<3; i++ {
			fmt.Print(". ")
			time.Sleep(1 * time.Second)
		}
		SlowPrint("\n\nT'as rien à me vendre? Vas t'en, je dois farm les primogems.")
		time.Sleep(3 * time.Second)
		p.Choixshops()
	}

	counter := 1
	var arrItem []string
	for item, nbr := range p.inventaire {
		fmt.Println(Colorize(Bold, Colorize(LightMagenta, strconv.Itoa(counter), " - "))+item+" : ", nbr)
		arrItem = append(arrItem, item)
		counter++
	}
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner aux shops. \n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	options := scanner.Text()
	number, err := strconv.Atoi(options)

	if err != nil {
		p.Collectionneur()
	}

	if number == 0 {
		p.Choixshops()
	}

	if number-1 < len(arrItem) {
		item := arrItem[number-1]
		p.Sellitem(item)
	} else {
		fmt.Println("Ecrire un choix existant.")
		time.Sleep(2 * time.Second)
		p.Collectionneur()
	}
}


// Sellitem permet de vendre au collectionneur un objet présent dans l'inventaire.
// L'argent est ajouté à celui du personnage et l'item supprimé de l'inventaire.
func (p *Personnage) Sellitem(item string) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as revendu :", item, "\n")
	if item == PtV || item == CdS || item == PdC{
		p.money++
		fmt.Println(Colorize(Yellow, "+1 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == PtM {
		p.money += 7
		fmt.Println(Colorize(Yellow, "+7 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == "Sort : Boule de feu" {
		p.money += 19
		fmt.Println(Colorize(Yellow, "+19 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == FdL {
		p.money += 2
		fmt.Println(Colorize(Yellow, "+4 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == TA || item == BA || item == CA || item == PtP || item == PdT {
		p.money += 4
		fmt.Println(Colorize(Yellow, "+4 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == PdG {
		p.money += 6
		fmt.Println(Colorize(Yellow, "+6 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == BdG {
		p.money += 8
		fmt.Println(Colorize(Yellow, "+8 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == HN || item == BN {
		p.money += 101
		fmt.Println(Colorize(Yellow, "+101 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == CD {
		p.money += 31
		fmt.Println(Colorize(Yellow, "+4 Coins"), "|| Tu as désormais ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	} else if item == DV {
		p.achatDV()
	} else {
		SlowPrint("T'essaies de m'arnaquer là? Vas t'en.")
		time.Sleep(3 * time.Second)
		p.Choixshops()
	}
	time.Sleep(3 * time.Second)
	p.RemoveInventory(item)
	p.Collectionneur()
}


// achatDV permet de vendre Davy au collectionneur.
// Possibilité de le garder ou le vendre.
func (p *Personnage) achatDV() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint("Sympa sympa... mais si tu veux t'en débarasser faudra me payer 31 Coins.\n\n")
	fmt.Print(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Vendre \n")
	fmt.Print(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Garder \n\n")
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "1":
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			SlowPrint("Bon choix.")
			p.money -= 31
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			fmt.Println("Tu t'es débarrassé de Davy.\n")
			fmt.Println(Colorize(Red, "-31 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
			time.Sleep(3 * time.Second)
			p.RemoveInventory(DV)
		case "2":
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			SlowPrint("Dommage...")
			time.Sleep(1 * time.Second)
		default:
			fmt.Print("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.achatDV()
		}
		p.Collectionneur()
	}
}