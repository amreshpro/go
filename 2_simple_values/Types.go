package main

import (
	"fmt"
)

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.999
	name       string  = "Amresh"
	intVar16   int16   = 3443
	intVar32   int32   = 2323
	intVar64   int64   = 43423
	intVar     int     = -3
	uintVar    uint    = 1
	uintVar32  uint32  = 12
	uintVar64  uint64  = 125
	uintVar8   uint8   = 0x1
	byteVar    byte    = 0x2
	runeVar    rune    = 'a' //only single quoe, double quote for string

)

// struct
type Player struct {
	name        string
	health      int
	attackPower float64
}

func getHealth(player Player) int {
	return player.health
}

func (player Player) getHealth() int {
	return player.health
}

//custom type
type Weapon string
func getWeapon(weapon Weapon) string {
	return string(weapon)
}


func Types() {

	player := Player{
		name:        "Amresh Maurya",
		health:      100,
		attackPower: 45.9,
	}
	fmt.Println(player)
	fmt.Printf("This is the player health: %d\n", getHealth(player))
	fmt.Printf("This is the player health: %d\n", player.health) //like a properties

	// map
	// users := map[string]int{} //empty map
	users := make(map[string]int)
	users["foo"] = 10
	users["bar"] = 11
	fmt.Printf("health:%+v\n", users)

	// for loop
	for key, value := range users {
		fmt.Printf("they key %s and the value %d \n", key, value)
	}

	age, ok := users["bar"]
	if !ok {
		fmt.Println("bar is not defined")
	} else {
		fmt.Println("age:", age)
	}

	delete(users, "bar") // only applicable for map , (map, key) for delete
	fmt.Println(users)

	// array
	numbers := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(numbers)

	// slices
	numbers2 := []int{}
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 10)
	numbers2 = append(numbers2, 45)

	fmt.Println(numbers2)

}
