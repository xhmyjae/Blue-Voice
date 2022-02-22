package main

import (
	"fmt"
	"time"
)


// SlowPrint permet de print une string lettre par lettre.
func SlowPrint(str ...string) {
	if DEBUG {
		fmt.Println(str)
	} else {
		for _, str_part := range str {
			for _, char := range str_part {
				fmt.Print(string(char))
				time.Sleep(40_000_000 * time.Nanosecond)
			}
		}
	}
}


const (
	Reset = "\033[0m"

	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Blue         = "\033[34m"
	Purple       = "\033[35m"
	Cyan         = "\033[36m"
	White        = "\033[37m"
	Gray         = "\033[90m"
	LightYellow  = "\033[93m"
	LightBlue    = "\033[94m"
	LightGreen   = "\033[92m"
	LightMagenta = "\033[95m"
	Bold         = "\033[1m"
	Dim          = "\033[2m"
	Italic       = "\033[3m"
	Underline    = "\033[4m"
)

// Colorize permet de changer la couleur d'une string.
func Colorize(color string, allStrings ...string) string {
	res := color
	for _, str := range allStrings {
		res += str
	}
	res += Reset
	return res
}


// Capitalize permet de mettre en majuscule la première lettre d'un mot.
func Capitalize(s string) string {
	var res string
	for index, cara := range s {
		if index != 0 {
			PInd := s[index-1]
			if !('a' <= PInd && PInd <= 'z' || '0' <= PInd && PInd <= '9' || 'A' <= PInd && PInd <= 'Z') {
				if 'a' <= cara && cara <= 'z' {
					cara -= 32
					res += string(cara)
				} else {
					res += string(cara)
				}
			} else {
				if 'A' <= cara && cara <= 'Z' {
					cara += 32
					res += string(cara)
				} else {
					res += string(cara)
				}
			}
		} else {
			if 'a' <= cara && cara <= 'z' {
				cara -= 32
				res += string(cara)
			} else {
				res += string(cara)
			}
		}
	}
	return res
}


// IsLetter vérifie si le caractère est une lettre.
func IsLetter(s string) bool {
	if len(s) == 0 {
		return true
	} else {
		for _, cara := range s {
			if !(rune('a') <= cara && cara <= rune('z') || rune('A') <= cara && cara <= rune('Z')) {
				return false
			}
		}
		return true
	}
}


// highlightTitle rajoute une "décoration" autour d'une string donnée.
func highlightTitle(str string) {
	fmt.Println("---------------", Colorize(Underline, str), "---------------")
}


// inList permet de vérifier si un array possède une certaine string.
func inList(str string, arr []string) bool {
	for _, each := range arr {
		if str == each {
			return true
		}
	}
	return false
}


// InvSize renvoie le nombre d'élément présent dans l'inventaire.
func InvSize(inv map[string]int) int {
	counter := 0
	for _, each := range inv {
		counter += each
	}
	return counter
}


// isEven revoie true si le nombre est pair.
func isEven(nbr int) bool {
	return nbr%2 != 1
}


// Noms de tous les items mis en variable.
const(
	PtV = "Potion de vie"
	PtP = "Potion de poison"
	PtM = "Potion de manaa"
	FdL = "Fourrure de loup"
	PdT = "Peau de troll"
	CdS = "Cuir de sanglier"
	PdC = "Plume de corbeau"
	CA = "Chapeau de l'aventurier"
	TA = "Tunique de l'aventurier"
	BA = "Bottes de l'aventurier"
	BP = "Bonnet de pauvre"
	MC = "Marcel"
	CC = "Claquettes avec chausettes"
	PdG = "Peau de Gobelin"
	BdG = "Bave de Gobelin"
	HN = "Haut pyjama Nami"
	BN = "Bas pyjama Nami"
	BPan = "Bonnet Panda"
	DV = "Davy"
	CD = "Chaîne Olivier"
	PF = "Peluche furry de Nami"
)
