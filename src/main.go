package main

import (
	"flag"
	"math/rand"
	"os"
	"time"
)

var DEBUG bool
var p1 Personnage
var c1 Chose
var c2 Chose
var c3 Chose
var c4 Chose
var date = 1

var firstnami = false

// C'est mieux de tester le jeu sur le terminal de vs code si possible, le build fait pas les effets attendus.


// main exécute le programme.
func main() {
	rand.Seed(time.Now().UnixNano())
	flag.BoolVar(&DEBUG, "debug", false, "run debug mode")
	flag.Parse()
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p1.CreatePerso()
	p1.Init(p1.nom, p1.classe, 1, p1.pdvmax, p1.pdvact, p1.initiative, p1.manaa, p1.manaamax, 0, 75, 8, 18, map[string]int{}, 10, p1.skill, 100, Equipement{helmet: BP, chestplate: MC, boots: CC})
	c1.InitChose("Cyril el gobelin", 40, 40, 8, 35, []string{PdG, BdG})
	c2.InitChose("Lory et Nami", 1000, 1000, 1000, 1000, []string{HN, BN})
	c3.InitChose("Olivier et son Davy", 120, 120, 25, 45, []string{DV, CD})
	c4.InitChose("Pierro", 300, 300, 134, 170, []string{PF})
	p1.Choixmenu()
}
