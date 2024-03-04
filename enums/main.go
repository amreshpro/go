package main

import "fmt"

//weapon type
//axe
//sword
//wooden stick
//knife

type WeaponType int

const (
	Axe         WeaponType = iota //increment -> 0
	Sword       //1
	WoodenStick //2
	Knife       //3
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 30
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Println("damage of the given weapon:", getDamage(Axe))

}
