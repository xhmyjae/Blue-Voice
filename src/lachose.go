package main

import (
	"strconv"
)

type Chose struct {
	nom    string
	pdvmax int
	pdvact int
	pdatt  int
	initiative int
	loot []string
}


// InitChose permet d'initier et faciliter la création du personnage.
func (c *Chose) InitChose(nom string, pdvmax int, pdvact int, pdatt int, initiative int, loot []string) {
	c.nom = nom
	c.pdvmax = pdvmax
	c.pdvact = pdvact
	c.pdatt = pdatt
	c.initiative = initiative
	c.loot = loot
}


// DisplayInitChose permet d'afficher les informations de la Chose.
func (c Chose) DisplayInitChose() []string {
	var res []string
	res = append(res, c.nom)
	res = append(res, Colorize(LightGreen, strconv.Itoa(c.pdvact), "/", strconv.Itoa(c.pdvmax)))
	return res
}
