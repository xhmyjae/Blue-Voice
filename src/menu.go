package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	//"time"
)


// Menu affiche l'interface des choix possibles dans le menu dont la possibilité de quitter le jeu.
func Menu() {
	highlightTitle("MENU")
	fmt.Println()
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Salle A7")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Shops")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 - ")), "Hautes Herbes")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "4 - ")), "Consulter les farming quests")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "5 - ")), "Histoire")
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour quitter le jeu.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}


// Choixmenu permet de sélectionner un des choix et le faire s'exécuter.
// Exécute la function Menu avant de permettre à l'utilisateur de faire des choix.
func (p *Personnage) Choixmenu() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	InFight = false
	InHistoire = false
	counterRound = 0
	Menu()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "1":
			p.Choixsalle()
		case "2":
			p.Choixshops()
		case "3":
			ChoixCombatPerso()
		case "4":
			p.Quests()
		case "5":
			p.CondHistoire()
		case "0":
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			os.Exit(0)
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.Choixmenu()
		}
	}
}
