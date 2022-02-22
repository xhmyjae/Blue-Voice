package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var FinHistoire = false
var checkpoint1 = false
var checkpoint2 = false
var checkpoint3 = false

var InHistoire bool


// CondHistoire vérifie si le joueur a le niveau suffisant pour faire le mode histoire.
func (p Personnage) CondHistoire() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if p.niveau < 45 {
		fmt.Println("Tu n'as pas le niveau requis pour vivre cette aventure...")
		time.Sleep(2 * time.Second)
		p.Choixmenu()
	} else {
		InHistoire = true
		InFight = true
		p.Histoire()
	}
}


// ChoixHistoire est une func de base permettant de faire des choix de dialogues.
func ChoixHistoire(a, b, c, d , e string) {
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 -")), a)
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 -")), b, "\n\n")
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
	Type := bufio.NewScanner(os.Stdin)
	Type.Scan()
	ChoixType := Type.Text()

	switch ChoixType {
	case "1":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Print(e)
		SlowPrint(c)
	case "2":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Print(e)
		SlowPrint(d)
	default:
		fmt.Println("Ecrire un choix existant.")
		time.Sleep(2 * time.Second)
		ChoixHistoire(a, b, c, d, e)
	}
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}


// Histoire lance l'histoire.
func (p *Personnage) Histoire() {
	SlowPrint(Colorize(Cyan, "Ca fait si longtemps qu'on marche... \n"))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("??? : ")
	SlowPrint("Attendez!")
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Mais t'es qui???"))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("Lory : ")
	SlowPrint("Je m'appelle Lory, et voici ma fidèle compagnon Nami! \n")
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("Lory : ")
	SlowPrint("J'aurais besoin de ton aide.. \n")
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Vas y dis nous tout."))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("Lory : ")
	SlowPrint(c4.nom, " a commis un acte irréparable... \n")
	SlowPrint("       Ses tendances étranges l'ont ammené à créer une attrocité... \n")
	time.Sleep(2 * time.Second)
	fmt.Println()

	ChoixHistoire("Qu'est-ce qu'il a bien pu faire de si grave?", "M'en bas les couilles frère", "I-Il a...", "Pas grave, tu vas quand même m'aider.", "Lory : ")

	fmt.Print("Lory : ")
	SlowPrint("Il a créé une peluche furry de Nami... \n")
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Imposible... Nous devons réagir!"))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Suivant les indications de Lory, la voix bleue et toi partirent à la recherche de", c1.nom)

	for i := 0; i<3; i++ {
		fmt.Print(" .")
		time.Sleep(1 * time.Second)
	}

	checkpoint1 = true
	village()
	Checkpoint1()
}


// village permet de faire une fois arriver dans le village.
func village() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Vous atteignez un village. Où souhaitez-vous aller?\n")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 -")), "Place du village")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 -")), "Maison chelou")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 -")), "Bar à gnomes\n")
	fmt.Print(Colorize(Bold, ">>> "))
	Type := bufio.NewScanner(os.Stdin)
	Type.Scan()
	ChoixType := Type.Text()

	switch ChoixType {
	case "1":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Un étranger te dit avoir vu", c1.nom, "dans la forêt.")
	case "2":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Une vieille sort de la maison, te regarde mal puis te claque la porte au nez.")
		time.Sleep(3 * time.Second)
		village()
	case "3":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tu repars du bar à gnomes aussi vite que tu es arrivé. Tu as gagné un trauma!")
		time.Sleep(3 * time.Second)
		village()
	default:
		village()
	}
}


// Checkpoint1 permet de faire revenir le personnage au 1er checkpoint.
func Checkpoint1() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Cette forêt est vraiment sombre..."))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Un cris rugit soudainement, vous comprenez que", c1.nom, "est là pour vous attaquer.")
	time.Sleep(4 * time.Second)

	mainfight(c1)
}


// Checkpoint2 permet de faire revenir le personnage au 2eme checkpoint.
func Checkpoint2() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Tu continues ton chemin et finis par sortir de la forêt.")
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Regarde! Un papillon! \n\n"))
	fmt.Println("Il s'agissait en réalité d'un pigeon parlant.")
	time.Sleep(4 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("Pigeon : ")
	SlowPrint("Je peux vous accompagner à ", c3.nom, " si vous le souhaitez.")

	ChoixHistoire("Oui", "Non", "Bon choix, suis moi le sang", "Trouve l'endroit seul bouffon.", "Pigeon : ")

	fmt.Println("C'est en continuant votre chemin que vous tombez nez-à-nez à ", c3.nom, " qui se baladaient.")
	time.Sleep(3 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint("Tu cherches ", c4.nom,"? Pas le choix, on va devoir se battre...")
	time.Sleep(2 * time.Second)

	mainfight(c3)
}


// Checkpoint3 permet de faire revenir le personnage au 3eme checkpoint.
func Checkpoint3() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Nous savons où se trouve", c4.nom, ". Allons y vite!"))
	time.Sleep(3 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Vous finissez par atteindre le lieu de vie de", c4.nom)
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("??? : ")
	SlowPrint("Montrez-vous sombre noob, je sais que vous êtes là.\n\n")
	fmt.Println("Forcés à vous montrez, tu te retrouves face à l'horrible", c4.nom)
	time.Sleep(3 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("Pierro : ")
	SlowPrint("Battons-nous, celui qui gagnera aura le droit de garder la peluche.\n\n")

	ChoixHistoire("J'accepte ton défis!", "Je n'accepte pas mais tu vas quand même me rendre la peluche!", "HAHA que le meilleur gagne..", "Viens te battre alors!", "Pierro : ")

	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Commence alors le combat qui changera tout l'avenir de Lory et de", p1.nom)
	time.Sleep(2 * time.Second)

	mainfight(c4)
}


// Checkpoint4 permet de faire revenir le personnage au 4eme checkpoint.
func Checkpoint4() {
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print("Lory : ")
	SlowPrint("Tu vois ", c4.nom, " tu ne pouvais pas battre le pouvoir de l'amitié qui me relie à Nami. \n\n")
	SlowPrint("Merci à toi ", p1.nom, " de m'avoir aidé! Je ne l'oublierais pas...")
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("Et c'est ainsi que fini la fabuleuse histoire d'un étranger qui avait rencontré Lory et Nami.")

	for i := 0; i<3; i++ {
		fmt.Print(" .")
		time.Sleep(1 * time.Second)
	}

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	SlowPrint(Colorize(Cyan, "Alors bête d'histoire pas vrai?? Enfin bon, continuons notre aventure."))
	time.Sleep(2 * time.Second)

	FinHistoire = true
	p1.quete3()
	InHistoire = false
	InFight = false
	p1.Choixmenu()
}
