package main

import "fmt"

type perso struct {
	name      string
	class     string
	lifemax   int
	life      int
	inventory []string
}

func (p *perso) init(name, class string, lifemax, life int, inventory []string) {
	p.name = name
	p.class = class
	p.lifemax = lifemax
	p.life = life
	p.inventory = inventory
}

func main() {
	var p1 perso
	p1.init("Sam", "Voleur", 100, 40, []string{"potion", "potion", "potion"})
	p1.displayinfo()
	p1.displayinventory()
	p1.takepot()
	p1.displayinfo()
}
func (p *perso) displayinfo() {
	fmt.Println("Nom:", p.name)
	fmt.Println("Classe:", p.class)
	fmt.Println("Vie max:", p.lifemax)
	fmt.Println("Vie actuel:", p.life)
	fmt.Println("Inventaire:", p.inventory)

}
func (p perso) displayinventory() {
	if p.inventory[0] == "" {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("Votre inventaire:")
		for i := 0; i <= len(p.inventory)-1; i++ {
			fmt.Println(p.inventory[i])
		}
	}
}
func (p *perso) takepot() {
	for i := 0; i <= len(p.inventory)-1; i++ {
		if p.inventory[i] == "potion" {
			p.life += 50
		} else {
			fmt.Println("Vous n'avez pas de potion")
		}
		if p.life > p.lifemax {
			p.life = p.lifemax
		}
		return
	}
}
