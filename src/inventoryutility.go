package main

import (
	"fmt"
	"time"
)


// useItem permet d'utiliser l'item.
// Chaque item à sa propre fonction d'exécution.
func (p *Personnage) useItem(item string) {
	if item == PtV {
		if p.pdvmax <= p.pdvact {
			fmt.Println("Tu es déjà au maximum de ta vie.")
		} else {
			p.ViePot()
			p.RemoveInventory(item)
		}
	} else if item == PtP {
		if InFight {
			if checkpoint1 {
				c1.ThrowPoisonPot()
			} else if checkpoint2 {
				c3.ThrowPoisonPot()
			} else if checkpoint3 {
				c4.ThrowPoisonPot()
			} else {
				c2.ThrowPoisonPot()
			}
		} else {
			p.PoisonPot()
		}
		p.RemoveInventory(item)
	} else if item == PtM {
		p.ManaaPot()
		p.RemoveInventory(item)
	} else if item == "Sort : Boule de feu" {
		p.AddSkills("Boule de feu")
		p.RemoveInventory(item)
	} else if item == CA || item == TA || item == BA|| item == BP || item == MC || item == CC || item == HN || item == BN || item == BPan {
		p.AddEquipement(item)
	} else {
		fmt.Println("Tu ne peux pas utiliser cet item ici.")
		time.Sleep(2 * time.Second)
		p.DisplayInventory()
	}
}


// AddInventory permet d'ajouter un item dans l'inventaire et retourne un bool.
// Si l'inventaire est rempli (au départ 10 items) il n'ajoute pas l'item et retourne false. Dans le cas où il peut, l'item est ajouté puis la
// function retourne true.
func (p *Personnage) AddInventory(item string) bool {
	if InvSize(p.inventaire) < p.limiteInv {
		if _, ok := p.inventaire[item]; ok {
			p.inventaire[item]++
			return true
		} else {
			p.inventaire[item] = 1
			return true
		}
	} else {
		fmt.Println("Ton inventaire est plein!")
		time.Sleep(2 * time.Second)
		return false
	}
}


// RemoveInventory permet de supprimer un item de l'inventaire.
// S'il n'y avait qu'un seul type de l'item dans l'inventaire, il sera supprimé, sinon son nombre diminuera de 1.
func (p *Personnage) RemoveInventory(item string) {
	if _, ok := p.inventaire[item]; ok {
		p.inventaire[item]--
		if p.inventaire[item] == 0 {
			delete(p.inventaire, item)
		}
	}
}
