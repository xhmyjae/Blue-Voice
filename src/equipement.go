package main

type Equipement struct {
	helmet     string
	chestplate string
	boots      string
}

var basepvhelmet int
var basepvchest int
var basepvboots int


// AddEquipement permet d'ajouter un des équipements craftés sur soi.
// Si jamais un item était déjà présent, il sera interchangé avec le nouveau puis replacer dans l'inventaire.
// La function fonctionne pour les 3 équipements actuellement présents.
func (p *Personnage) AddEquipement(armor string) {
	var res string
	allhelmet := []string{CA, BP, BPan}
	allchest := []string{TA, MC, HN}
	allboots := []string{BA, CC, BN}
	valhelmet := []int{10, 0, 100}
	valchest := []int{25, 0, 100}
	valboots := []int{15, 0, 100}

	for ind, val := range allhelmet {
		if armor == val {
			if len(p.equipement.helmet) != 0 {
				for _, each := range p.equipement.helmet {
					res += string(each)
				}
				p.AddInventory(res)
			}
			p.equipement.helmet = allhelmet[ind]
			p.pdvmax -= basepvhelmet
			basepvhelmet = valhelmet[ind]
			p.pdvmax += basepvhelmet
		}
	}

	for ind, val := range allchest {
		if armor == val {
			if len(p.equipement.chestplate) != 0 {
				for _, each := range p.equipement.chestplate {
					res += string(each)
				}
				p.AddInventory(res)
			}
			p.equipement.chestplate = allchest[ind]
			p.pdvmax -= basepvchest
			basepvchest = valchest[ind]
			p.pdvmax += basepvchest
		}
	}

	for ind, val := range allboots {
		if armor == val {
			if len(p.equipement.boots) != 0 {
				for _, each := range p.equipement.boots {
					res += string(each)
				}
				p.AddInventory(res)
			}
			p.equipement.boots = allboots[ind]
			p.pdvmax -= basepvboots
			basepvboots = valboots[ind]
			p.pdvmax += basepvboots
		}
	}
	p.RemoveInventory(armor)

	if !(firstnami) {
		p.quete5()
	}
}
