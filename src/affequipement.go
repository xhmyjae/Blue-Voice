package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


// Chapeau permet de crafter le "helmet" du personnage.
// Si tous les éléments requis sont présents, le helmet sera créé et ajouté dans l'inventaire (les matériaux en seront eux supprimés).
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du forgeron (cf. func Choixforgeron).
func (p *Personnage) Chapeau() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if p.inventaire[PdC] < 1 && p.inventaire[CdS] < 1 {
		fmt.Println("Tu n'as pas les objects requis : Plume de corbeau (1), Cuir de sanglier (1).")
	} else if p.inventaire[PdC] >= 1 && p.inventaire[CdS] < 1 {
		fmt.Println("Tu n'as pas les objects requis : Cuir de sanglier (1).")
	} else if p.inventaire[PdC] < 1 && p.inventaire[CdS] >= 1 {
		fmt.Println("Tu n'as pas les objects requis : Plume de corbeau (1).")
	}
	if p.inventaire[PdC] >= 1 && p.inventaire[CdS] >= 1 {
		if p.AddInventory(CA) {
			p.RemoveInventory(PdC)
			p.RemoveInventory(CdS)
			fmt.Println("Tu as crafté le Chapeau de l'aventurier! Retrouve le dans votre inventaire.\n")
			p.money -= 5
			fmt.Println(Colorize(Red, "-5 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		}
	}
	time.Sleep(2 * time.Second)
	p.Choixforgeron()
}


// Tunique permet de crafter le "chestplate" du personnage.
// Si tous les éléments requis sont présents, le chestplate sera créé et ajouté dans l'inventaire (les matériaux en seront eux supprimés).
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du forgeron (cf. func Choixforgeron).
func (p *Personnage) Tunique() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if p.inventaire[FdL] < 1 && p.inventaire[PdT] < 2 {
		fmt.Println("Tu n'as pas les objects requis : Fourrure de loup (1), Peau de troll (2).")
	} else if p.inventaire[FdL] >= 1 && p.inventaire[PdT] < 2 {
		fmt.Println("Tu n'as pas les objects requis : Peau de troll (2).")
	} else if p.inventaire[FdL] < 1 && p.inventaire[PdT] >= 2 {
		fmt.Println("Tu n'as pas les objects requis : Fourrure de loup (1).")
	}
	if p.inventaire[FdL] >= 1 && p.inventaire[PdT] >= 2 {
		if p.AddInventory(TA) {
			p.RemoveInventory(FdL)
			p.RemoveInventory(PdT)
			p.RemoveInventory(PdT)
			fmt.Println("Tu as crafté la Tunique de l'aventurier! Retrouve la dans ton inventaire.\n")
			p.money -=5
			fmt.Println(Colorize(Red, "-5 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		}
	}
	time.Sleep(2 * time.Second)
	p.Choixforgeron()
}


// Bottes permet de crafter les "boots" du personnage.
// Si tous les éléments requis sont présents, les boots seront créées et ajoutées dans l'inventaire (les matériaux en seront eux supprimés).
// Une fois les messages correspondants affichés, l'utilisateur sera redirigé à la page Choix du forgeron (cf. func ChoixForgeron).
func (p *Personnage) Bottes() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if p.inventaire[FdL] < 1 && p.inventaire[CdS] < 1 {
		fmt.Println("Tu n'as pas les objects requis : Fourrure de loup (1), Cuir de sanglier (1).")
	} else if p.inventaire[FdL] >= 1 && p.inventaire[CdS] < 1 {
		fmt.Println("Tu n'as pas les objects requis : Cuir de sanglier (1).")
	} else if p.inventaire[FdL] < 1 && p.inventaire[CdS] >= 1 {
		fmt.Println("Tu n'as pas les objects requis : Fourrure de loup (1).")
	}
	if p.inventaire[FdL] >= 1 && p.inventaire[CdS] >= 1 {
		if p.AddInventory(BA) {
			p.RemoveInventory(FdL)
			p.RemoveInventory(CdS)
			fmt.Println("Tu as crafté les Bottes de l'aventurier! Retrouve les dans ton inventaire.\n")
			p.money -= 5
			fmt.Println(Colorize(Red, "-5 Coins"), "|| Il te reste ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
		}
	}
	time.Sleep(2 * time.Second)
	p.Choixforgeron()
}
