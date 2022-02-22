package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var counterRound int
var InFight bool


// ChoixCombatPerso permet de lancer le combat contre la chose.
// Le mob sera choisi aléatoirement (+ de chance pour c1 que c3 que c2)
// Le personnage avec le plus d'initiative commence le combat.
func ChoixCombatPerso() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	randnb:= rand.Intn(9)+1

	var mob Chose

	if randnb == 1 || randnb == 2 || randnb == 3 || randnb == 4 {
		mob = c1
	} else if randnb == 5 || randnb == 6 || randnb == 7 || randnb == 8{
		mob = c3
	} else {
		mob = c2
	}

	InFight = true
	fmt.Println(mob.nom, "veut se battre contre vous!\n")
	SlowPrint(Colorize(Cyan, "Ca va être rude pour toi! Je reste à tes côtés t'inquiètes!"))
	time.Sleep(2 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("La voix bleue s'en va et te laisse seul contre", mob.nom)
	time.Sleep(3 * time.Second)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	mob.DisplayInitChose()

	if mob.initiative > p1.initiative {
		fmt.Println(mob.nom, "commence le combat!")
		time.Sleep(2 * time.Second)
		p1.ChoseAttaque(mob)
	} else {
		p1.PersoAttaque(mob)
	}
}


// ChoseAttaque permet à la Chose de nous attaquer en faisant des dégats critiques tous les 3 tours.
func (p *Personnage) ChoseAttaque(mob Chose) {
	counterRound++
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("C'est au tour de", mob.nom, "!\n")

	if mob.nom == c1.nom {
		fmt.Println(mob.nom, "utilise Calvicie sur toi!")
		fmt.Println(Colorize(Cyan, "Oh non! Tu commences à développer une calvicie!"))
	} else if mob.nom == c2.nom {
		fmt.Println(mob.nom, "utilise God Nami sur toi, sa puissante lumière t'aveugle.")
	} else if mob.nom == c4.nom {
		fmt.Println(mob.nom, "utilise Code Minecraft sur toi! Tu manques de te changer en cube!")
	} else {
		fmt.Println(mob.nom, "utilise racisme et envoie Davy t'attaquer.")
	}
	time.Sleep(2 * time.Second)

	if counterRound%3 == 0 {
		dmgdouble := mob.pdatt*2
		p.pdvact -= dmgdouble
		if p.pdvact > 0 {
			fmt.Println("Cependant", mob.nom, "t'infliges un dégat critique!")
			fmt.Println("Tu perds donc ", Colorize(Red, strconv.Itoa(dmgdouble), "PV"), "|| Il te reste ", Colorize(LightGreen, strconv.Itoa(p.pdvact), "/", strconv.Itoa(p.pdvmax), "PV"))
			time.Sleep(3 * time.Second)
			if InHistoire {
				p.P1Attaque(mob)
			} else {
				p.PersoAttaque(mob)
			}
		} else {
			if InHistoire {
				p.DeadHistoire()
			} else {
				p.Dead()
			}
			time.Sleep(3 * time.Second)
			p.Choixmenu()
		}
	}
	p.pdvact -= mob.pdatt
	fmt.Println()

	if p.pdvact > 0 {
		fmt.Println("Tu perds ", Colorize(Red, strconv.Itoa(mob.pdatt), "PV"), "|| Il te reste ", Colorize(LightGreen, strconv.Itoa(p.pdvact), "/", strconv.Itoa(p.pdvmax), "PV"))
		time.Sleep(3 * time.Second)
		if InHistoire {
			p.P1Attaque(mob)
		} else {
			p.PersoAttaque(mob)
		}
	} else {
		if InHistoire {
			p.DeadHistoire()
		} else {
			p.Dead()
		}
		time.Sleep(3 * time.Second)
		p.Choixmenu()
	}
}


// PersoAttaque permet au personnage d'attaquer, utiliser l'inventaire ou fuir.
func (p *Personnage) PersoAttaque(mob Chose) {
	counterRound++
	Choixcombat()
	for {
		scannerFight := bufio.NewScanner(os.Stdin)
		scannerFight.Scan()
		optionsFight := scannerFight.Text()
		switch optionsFight {
		case "1":
			p.SkillUse(mob)
		case "2":
			p.invCombat(mob)
		case "3":
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			SlowPrint(Colorize(Cyan, "Tu fuis le combat? Pfff... Si nul."))
			time.Sleep(1 * time.Second)
			fmt.Println()
			SlowPrint(Colorize(Cyan, " Je te comprends en vrai, partons vite"))
			time.Sleep(2 * time.Second)
			mob.pdvact = mob.pdvmax
			p.Choixmenu()
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.PersoAttaque(mob)
		}
	}
}


// Choixcombat affiche les choix possibles par le personnage.
func Choixcombat() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("C'est ton tour!\n")
	fmt.Println("Tes manaas actuels : ", Colorize(LightBlue, strconv.Itoa(p1.manaa)), "\n")
	fmt.Println("Que décides-tu de faire?\n")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "1 - ")), "Compétence")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Inventaire")
	fmt.Println(Colorize(Bold, Colorize(LightMagenta, "3 - ")), "Fuir\n\n")
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))
}


// SkillUse permet d'utiliser un skill.
// Si le joueur a acheté le skill boule de feu, il sera aussi disponible parmi les skills utilisables.
func (p *Personnage) SkillUse(mob Chose) {
	num:= rand.Intn(9)+1

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Print(Colorize(Bold, Colorize(LightMagenta, "1 - ")), p.skill[0])
	fmt.Println()
	for _, each := range p.skill {
		if each == "Boule de feu" {
			fmt.Print(Colorize(Bold, Colorize(LightMagenta, "2 - ")), "Boule de feu")
			fmt.Println()
		}
	}
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour revenir au menu précédent.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))

	allskills := []string{"Attaque basique", "Jugement de l'elfe", "Lancé de hachette", "Pitzah"}
	allfails := []string{"Tu glisses en essayant de frapper l'ennemi, aucun effort...", "Tu n'étais pas aussi beau que tu le pensais, ton ego en prends un coup.", "Tu viens de faire tomber la hachette sur ton pied... Ton cerveau est aussi petit que toi.", "Tu as mis de l'ananas sur la pizza! Tu ne mérites pas d'être italien."}
	allw := []string{"Tu infliges un coup de poing basique car tu n'as pas la capacité de faire mieux!", "Tu blesses son ego en l'exposant à un miroir.", "Tu n'as jamais réussi à bien viser, mais cette fois-ci ça a marché!", "Etant italien, tu as des pizzas naturellement sur toi et en lance une sur l'ennemi."}

	for {
		scannerSkill := bufio.NewScanner(os.Stdin)
		scannerSkill.Scan()
		optionsSkill := scannerSkill.Text()

		switch optionsSkill {
		case "1":
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			if p.manaa >= 5 {
				p.manaa -= 5
				for ind, val := range allskills {
					if p.skill[0] == val {
						if num == 1 {
							fmt.Println(allfails[ind])
							time.Sleep(3 * time.Second)
							p.pdvact -= p.damage
							if p.pdvact <= 0 {
								if InHistoire {
									p.DeadHistoire()
								} else {
									p.Dead()
								}
								time.Sleep(3 * time.Second)
								p.Choixmenu()
							} else {
								fmt.Println("Tu perds ", Colorize(Red, strconv.Itoa(p.damage), "PV"), "|| Il te reste ", Colorize(LightGreen, strconv.Itoa(p.pdvact), "/", strconv.Itoa(p.pdvmax), "PV"))
								time.Sleep(3 * time.Second)
								p.ChoseAttaque(mob)
							}
						} else {
							mob.pdvact -= p.damage
							fmt.Println(allw[ind])
							time.Sleep(3 * time.Second)
							if mob.pdvact > 0 {
								fmt.Println("\n", mob.nom, "perd", strconv.Itoa(p.damage), "PV! Il lui en reste ", Colorize(Gray, strconv.Itoa(mob.pdvact), "/", strconv.Itoa(mob.pdvmax)))
								time.Sleep(4 * time.Second)
								p.ChoseAttaque(mob)
							} else {
								if InHistoire {
									mob.ChoseDeadHistoire()
									break
								}
								mob.ChoseDead()
							}
						}
					}
				}
			} else {
				fmt.Println("Tu n'as pas assez de manaa.")
				time.Sleep(2 * time.Second)
				Choixcombat()
			}
		case "2":
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			if p.manaa >= 15 {
				p.manaa -= 15
				fmt.Println("Tu envoies des boules de feu partout!")
				time.Sleep(3 * time.Second)
				if  num == 1 {
					fmt.Println("Tes boules de feu se retournent contre toi! Apprends à discerner ta personne de ton adversaire. Sale noob.")
					time.Sleep(3 * time.Second)
					p.pdvact -= p.damageboule
					if p.pdvact <= 0 {
						if InHistoire {
							p.DeadHistoire()
						}
						p.Dead()
						time.Sleep(3 * time.Second)
						p.Choixmenu()
					} else {
						fmt.Println(mob.nom, "perd", strconv.Itoa(p.damageboule), "PV! Il lui en reste ", Colorize(Gray, strconv.Itoa(mob.pdvact), "/", strconv.Itoa(mob.pdvmax)))
						time.Sleep(3 * time.Second)
						p.ChoseAttaque(mob)
					}
				} else {
					mob.pdvact -= p.damageboule
					if mob.pdvact > 0 {
						fmt.Println(mob.nom, "s'en reçoit une et perd", strconv.Itoa(p.damageboule),"PV! Il lui en reste ", mob.pdvact, "/", mob.pdvmax)
						time.Sleep(3 * time.Second)
						p.ChoseAttaque(mob)
					} else {
						if InHistoire {
							mob.ChoseDeadHistoire()
							break
						}
						mob.ChoseDead()
					}
				}
			} else {
				fmt.Println("Tu n'as pas assez de manaa.")
				time.Sleep(2 * time.Second)
				Choixcombat()
			}
		case "0":
			p.PersoAttaque(mob)
		default:
			fmt.Println("Ecrire un choix existant.")
			time.Sleep(2 * time.Second)
			p.SkillUse(mob)
		}
	}
}


// invCombat permet d'afficher et utiliser l'inventaire durant le combat.
func (p * Personnage) invCombat(mob Chose) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	highlightTitle("Inventaire")
	fmt.Println()
	if len(p.inventaire) == 0 {
		fmt.Println(Colorize(Italic, "Ton inventaire est vide."))
		time.Sleep(2 * time.Second)
		if InHistoire {
			p.P1Attaque(mob)
		} else {
			p.PersoAttaque(mob)
		}
	}

	counter := 1
	var arrItem []string
	for item, nbr := range p.inventaire {
		fmt.Println(Colorize(Bold, Colorize(LightMagenta, strconv.Itoa(counter), "- "))+item+" : ", nbr)
		arrItem = append(arrItem, item)
		counter++
	}
	fmt.Println(Colorize(Purple, "Appuie sur 0 pour retourner au menu précédent.\n\n"))
	fmt.Println("Entre ton choix : ")
	fmt.Print(Colorize(Bold, ">>> "))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	options := scanner.Text()
	number, err := strconv.Atoi(options)

	if err != nil {
		p.invCombat(mob)
	}

	if number == 0 {
		if InHistoire {
			p.P1Attaque(mob)
		} else {
			p.PersoAttaque(mob)
		}
	}

	if number-1 < len(arrItem) {
		item := arrItem[number-1]
		p.useItem(item)
		time.Sleep(2 * time.Second)
		p.ChoseAttaque(mob)
	} else {
		fmt.Println("Ecrire un choix existant.")
		time.Sleep(2 * time.Second)
		p.invCombat(mob)
	}
}
