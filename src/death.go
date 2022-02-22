package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)


// Dead fait mourir le personnage si jamais ses PVs tombent à 0.
// À la fin de la func, le personnage sera réanimé avec la moitié de ses PV et Manaa max, puis redirigé vers le menu.
func (p *Personnage) Dead() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if p.pdvact <= 0 {
		time.Sleep(1 * time.Second)
		fmt.Println("Tu as atteint ta limite!")
		fmt.Println(Colorize(Red, "0/", strconv.Itoa(p.pdvmax), "PV\n"))
		time.Sleep(3 * time.Second)
		SlowPrint(Colorize(Cyan, "HAHA Noob you died "))
		for i := 0; i<3; i++ {
			fmt.Print(". ")
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
		SlowPrint(Colorize(Cyan, "Mais vas-y reviens fréro, je te laisse une autre chance."))
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.pdvact = p.pdvmax / 2
		p.manaa = p.manaamax / 2
		fmt.Println("Tu as été réanimé par la voix bleue.")
	}
}


// ChoseDead fait mourir la Chose si jamais ses PV tombent à 0.
// Le Personnage gagne de l'expérience à ce moment.
// Différentes versions sont disponibles selon le mob tué.
// Les PV des mobs reviennent à leur max à la fin du combat.
func (c *Chose) ChoseDead() {
	Indnum := rand.Intn(7)+1

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as vaincu", c.nom)
	time.Sleep(2 * time.Second)
	p1.GainExp(50 )

	if Indnum == 1 || Indnum == 2 || Indnum == 3 || Indnum == 4 {
		p1.AddInventory(c.loot[0])
		p1.money += Indnum-1
		if c.nom == c3.nom {
			fmt.Print("Davy : ")
			SlowPrint("Vous êtes largement plus fort que mon maitre, laissez moi vous rejoindre! \n\n")
			fmt.Println("Davy te rejoint et te donne", Colorize(LightYellow, strconv.Itoa(Indnum-1), " Coins"), " par la même occasion.")
		} else if c.nom == c2.nom {
			fmt.Print("Lory : ")
			SlowPrint("Je reconnais ton talent, tiens prends ce haut de pyjamas Nami. \n\n")
			fmt.Println("Lory te donne le haut de pyjama Nami! Equipe depuis ton inventaire. Elle te donne aussi", Colorize(LightYellow, strconv.Itoa(Indnum-1), " Coins"))
			p1.quete4()
		} else {
			fmt.Print("Cyril : ")
			SlowPrint("Ma calvicie ne t'aura donc pas atteint... \n\n")
			fmt.Println(c.nom, "drop", Colorize(LightYellow, strconv.Itoa(Indnum-1), " Coins"), ", et", c.loot[0])
			p1.quete1()
		}
	} else {
		p1.AddInventory(c.loot[1])
		p1.money += Indnum-1
		if c.nom == c3.nom {
			fmt.Print("Davy : ")
			SlowPrint("Merci d'avoir battu mon maître.. je laisse tomber mes chaînes pour vous. \n\n")
			fmt.Println("Davy t'as donné ses chaînes! Il te donne", Colorize(LightYellow, strconv.Itoa(Indnum-1), " Coins"), " par la même occasion.")
		} else if c.nom == c2.nom {
			fmt.Print("Lory : ")
			SlowPrint("Je reconnais ton talent, tiens prends ce bas de pyjamas Nami. \n\n")
			fmt.Println("Lory te donne le bas de pyjama Nami! Equipe depuis ton inventaire. Elle te donne", Colorize(LightYellow, strconv.Itoa(Indnum-1), " Coins"), " par la même occasion.")
			p1.quete4()
		} else {
			fmt.Print("Cyril : ")
			SlowPrint("Ma calvicie ne t'aura donc pas atteint... \n\n")
			fmt.Println(c.nom, "drop", Colorize(LightYellow, strconv.Itoa(Indnum-1), " Coins"), ", et", c.loot[1])
			p1.quete1()
		}
	}
	time.Sleep(3 * time.Second)

	c.pdvact = c.pdvmax
	p1.Choixmenu()
}


// DeadHistoire fait mourir le personnage quand ses PV tombent à 0 dans le mode histoire.
// Il sera ranimé avec 50% de ses PV à l'un des checkpoints, sinon il peut choisir de retourner au menu.
func (p *Personnage) DeadHistoire() {
	if p.pdvact <= 0 {
		time.Sleep(1 * time.Second)
		fmt.Println("Tu n'étais pas à la hauteur...")
		fmt.Println(Colorize("0/", strconv.Itoa(p.pdvmax), " PV\n"))
		time.Sleep(3 * time.Second)
		SlowPrint(Colorize(Cyan, "HAHA Noob you died"))
		fmt.Print(" ")
		for i := 0; i<3; i++ {
			fmt.Print(". ")
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.pdvact = p.pdvmax / 2
		ChoixMort()
	}
}


// ChoseDeadHistoire fait mourir les mobs dans l'histoire.
// Le mob reprendra ses PV à la fin, mais le joueur ne gagnera pas d'exp dans ce mode.
func (c *Chose) ChoseDeadHistoire() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu as vaincu", c.nom, "\n")
	if c.nom == c1.nom {
		fmt.Print(c.nom, ": ")
		SlowPrint("Tu m'as peut-être battu.. mais jamais tu ne vaincras", c3.nom, "!\n\n")
		SlowPrint(Colorize(Cyan, "Mais il ne faut pas oublier notre objectif de base. Continuons!"))
		time.Sleep(3 * time.Second)
		checkpoint1 = false
		checkpoint2 = true
		Checkpoint2()
	}
	if c.nom == c3.nom {
		fmt.Print("Olivier : ")
		SlowPrint("Je te dirais tout ce que tu veux savoir mais je t'en pris, ne touche pas à mon Davy... \n\n")
		SlowPrint(Colorize(Cyan, "Mais il ne faut pas oublier notre objectif de base. Continuons!"))
		time.Sleep(3 * time.Second)
		checkpoint2 = false
		checkpoint3 = true
		Checkpoint3()
	}
	if c.nom == c4.nom {
		p1.AddInventory(c.loot[0])
		fmt.Print("Pierro : ")
		SlowPrint("Je n'aurais jamais cru que tu aurais pu me vaincre... \n\n")
		fmt.Println(c4.nom, "rends la peluche furry à Lory.")
		time.Sleep(3 * time.Second)
		checkpoint3 = false
		Checkpoint4()
	}
	c.pdvact = c.pdvmax
	time.Sleep(3 * time.Second)
}


// ChoixMort permet à l'utilisateur de choisir s'il veut continuer ou s'arrêter de jouer après sa mort dans le mode histoire.
func ChoixMort() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Souhaites-tu continuer le mode histoire?")
	fmt.Println()
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 -")), "Continuer")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 -")), "Arrêter")
	fmt.Println()
	fmt.Println()
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
	Type := bufio.NewScanner(os.Stdin)
	Type.Scan()
	ChoixType := Type.Text()

	switch ChoixType {
	case "1":
		SlowPrint(Colorize(Cyan, "Aller on y retourne!"))
		time.Sleep(2 * time.Second)
		if checkpoint1 {
			Checkpoint1()
		}
		if checkpoint2 {
			Checkpoint2()
		}
		if checkpoint3 {
			Checkpoint3()
		}
	case "2":
		SlowPrint(Colorize(Cyan, "Mais t'inquiètes fréro, on retournera aider Lory une autre fois."))
		time.Sleep(2 * time.Second)
		p1.Choixmenu()
	default:
		fmt.Println("Ecrire un choix existant.")
		time.Sleep(2 * time.Second)
		ChoixMort()
	}
}