package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Personnage struct {
	nom        	string
	classe     	string
	niveau    	int
	pdvmax    	int
	pdvact    	int
	initiative	int
	manaa 		int
	manaamax 	int
	exp 		int
	expmax 		int
	damage		int
	damageboule int
	inventaire 	map[string]int
	limiteInv  	int
	skill      	[]string
	money      	int
	equipement 	Equipement
}


// Init initialise la création du personnage.
// Elle utilise la struct Personnage créée auparavant.
func (p *Personnage) Init(nom string, classe string, niveau int, pdvmax int, pdvact int, initiative int, manaa int, manaamax int, exp int, expmax int, damage int, damageboule int, inventaire map[string]int, limiteInv int, skill []string, money int, equipement Equipement) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.exp = exp
	p.expmax = expmax
	p.damage = damage
	p.damageboule = damageboule
	p.pdvmax = pdvmax
	p.pdvact = pdvact
	p.initiative = initiative
	p.manaa = manaa
	p.manaamax = manaamax
	p.inventaire = inventaire
	p.limiteInv = limiteInv
	p.skill = skill
	p.money = money
	p.equipement = equipement
}


// CreatePerso permet à l'utilisateur de créer son personnage au démarrage du jeu.
func (p *Personnage) CreatePerso() {
	SlowPrint(Colorize(Cyan, "Bienvenue dans Blue Voice!"), "\n\n")
	SlowPrint(Colorize(Cyan, "J'espère que tu te plairas ici, c'est plutôt fun tu verras..."))
	time.Sleep(1 * time.Second)

	p.Choixnom()

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, p.nom, "... Je vois. J'aime bien, ça te va bien!"))
	time.Sleep(1 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Sinon ", p.nom, ", dis m'en un peu plus sur toi.\n"))
	fmt.Println()

	p.Choixrace()

	SlowPrint(Colorize(Cyan, "Bon, mon reuf.. il est temps de continuer ton aventure. Laisses moi t'accompagner."))
	time.Sleep(1 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("La voix bleue est partie")
	fmt.Print(" ")

	for i := 0; i<3; i++ {
		fmt.Print(". ")
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}


// Choixnom permet à l'utilisateur de mettre son nom, si celui-ci est nil alors la func se relance.
func (p *Personnage) Choixnom() string {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Dis moi, c'est quoi ton ptit nom?\n\n"))
	SlowPrint("Entre ton nom :", "\n")
	fmt.Print(Colorize(Bold, ">>> "))

	INPersoNom := bufio.NewScanner(os.Stdin)
	INPersoNom.Scan()
	PersoNom := INPersoNom.Text()

	if !IsLetter(PersoNom){
		fmt.Println("Ne mettre que des lettres.")
		time.Sleep(2 * time.Second)
		p.Choixnom()
	} else {
		p.nom = Capitalize(PersoNom)
	}
	return p.nom
}


// Choixrace permet à l'utilisateur de choisir la classe de son personnage.
func (p *Personnage) Choixrace() {
	SlowPrint("Choisis la race de ton personnage en écrivant le chiffre associé :\n\n")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 -")), " Humain")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 -")), " Elfe")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 -")), " Nain\n\n")
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))

	INRaceType := bufio.NewScanner(os.Stdin)
	INRaceType.Scan()
	RaceType := INRaceType.Text()

	switch RaceType {
	case "1":
		p.classe = "Humain"
		p.pdvmax = 100
		p.pdvact = 50
		p.initiative = 20
		p.manaa = 35
		p.manaamax = 85
		p.skill = append(p.skill, "Attaque basique")
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		SlowPrint(Colorize(Cyan, "Un humain! C'est bien ce que je me disais. T'es vraiment lambda."))
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	case "2":
		p.classe = "Elfe"
		p.pdvmax = 80
		p.pdvact = 40
		p.initiative = 30
		p.manaa = 60
		p.manaamax = 120
		p.skill = append(p.skill, "Jugement de l'elfe")
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		SlowPrint(Colorize(Cyan, "Un elfe! Si tu voulais être plus original t'aurais pu prendre humain."))
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	case "3":
		p.classe = "Nain"
		p.pdvmax = 120
		p.pdvact = 60
		p.initiative = 80
		p.manaa = 30
		p.manaamax = 70
		p.skill = append(p.skill, "Lancé de hachette")
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		SlowPrint(Colorize(Cyan, "Un nain! Dur, je te soutiens mon pote..."))
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	default:
		p.classe = "Italien"
		p.pdvmax = 50
		p.pdvact = 25
		p.initiative = 45
		p.manaa = 30
		p.manaamax = 80
		p.skill = append(p.skill, "Pitzah")
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		SlowPrint(Colorize(Cyan, "T'as pas bien compris la consigne, tu dois être un italien."))
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	}
}


// DisplayInit affiche les détails du personnage créé.
func (p *Personnage) DisplayInit() {
	highlightTitle("Informations du personnage")
	fmt.Println()
	fmt.Println("Nom : " + p.nom)
	fmt.Println("Race : " + p.classe)
	fmt.Println("Level : ", p.niveau)
	fmt.Println("Expérience : ", strconv.Itoa(p.exp), "/", strconv.Itoa(p.expmax))
	fmt.Println("PV : ", Colorize(LightGreen, strconv.Itoa(p.pdvact), "/", strconv.Itoa(p.pdvmax)))
	fmt.Println("Manaa : ", Colorize(LightBlue, strconv.Itoa(p.manaa), "/", strconv.Itoa(p.manaamax)))
	fmt.Println("Skills : ", strings.Join(p.skill, ", "))
	fmt.Println("Argent : ", Colorize(LightYellow, strconv.Itoa(p.money), " Coins"))
	fmt.Println("Equipement : ", "\n  - Casque: ", p.equipement.helmet, "\n  - Plastron: ", p.equipement.chestplate, "\n  - Bottes: ", p.equipement.boots)
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner au menu.\n\n"))
}


// ChoixInfo permet à l'utilisateur de retourner au menu s'il le souhaite après avoir affiché les informations du personnage.
func (p *Personnage) ChoixInfo() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.DisplayInit()
	fmt.Print(Colorize(Bold, ">>> "))

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		options := scanner.Text()

		switch options {
		case "0":
			p.Choixsalle()
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.ChoixInfo()
		}
	}
}


// GainExp permet de faire gagner de l'exp au personnage et augmenter le niveau du Personnage.
func (p * Personnage) GainExp(exp int) {
	p.exp += exp
	if p.exp >= p.expmax {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		diffexp := p.exp - p.expmax
		p.niveau++
		p.damage += 3
		p.damageboule += 2
		p.pdvmax++
		p.initiative++
		p.exp = 0
		fmt.Println("Gain d'expérience! Tu es désormais au niveau", p.niveau)
		time.Sleep(2 * time.Second)
		p.GainExp(diffexp)
	}
	p.quete2()
}
