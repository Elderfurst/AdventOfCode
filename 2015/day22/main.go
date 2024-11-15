package main

import (
	"adventofcode/utilities"
	"fmt"
	"math"
)

func main() {
	boss := fighter{
		hitpoints: 55,
		damage:    8,
		armor:     0,
	}

	shop := buildShop()

	ringSubsets := utilities.GetSubsets(shop.rings)

	ringOptions := make([][]item, 0)

	for _, subset := range ringSubsets {
		if len(subset) <= 2 {
			ringOptions = append(ringOptions, subset)
		}
	}

	lowestCost := math.MaxInt
	highestCost := math.MinInt

	for i := 0; i < len(shop.weapons); i++ {
		for j := -1; j < len(shop.armor); j++ {
			for k := 0; k < len(ringOptions); k++ {
				cost := 0

				player := fighter{
					hitpoints: 100,
					damage:    0,
					armor:     0,
				}

				player.damage += shop.weapons[i].damage
				cost += shop.weapons[i].cost

				if j >= 0 {
					player.armor += shop.armor[j].armor
					cost += shop.armor[j].cost
				}

				for _, ring := range ringOptions[k] {
					player.damage += ring.damage
					player.armor += ring.armor

					cost += ring.cost
				}

				if fight(player, boss) {
					if cost < lowestCost {
						lowestCost = cost
					}
				} else {
					if cost > highestCost {
						highestCost = cost
					}
				}
			}
		}
	}

	fmt.Println(lowestCost)
	fmt.Println(highestCost)
}

func fight(player, boss fighter) bool {
	for {
		bossDamage := utilities.Max(1, player.damage-boss.armor)

		boss.hitpoints -= bossDamage

		if boss.hitpoints <= 0 {
			return true
		}

		playerDamage := utilities.Max(1, boss.damage-player.armor)

		player.hitpoints -= playerDamage

		if player.hitpoints <= 0 {
			return false
		}
	}
}

func buildShop() shop {
	weapons := []item{
		{name: "Dagger", cost: 8, damage: 4, armor: 0},
		{name: "Shortsword", cost: 10, damage: 5, armor: 0},
		{name: "Warhammer", cost: 25, damage: 6, armor: 0},
		{name: "Longsword", cost: 40, damage: 7, armor: 0},
		{name: "Greataxe", cost: 74, damage: 8, armor: 0},
	}

	armor := []item{
		{name: "Leather", cost: 13, damage: 0, armor: 1},
		{name: "Chainmail", cost: 31, damage: 0, armor: 2},
		{name: "Splintmail", cost: 53, damage: 0, armor: 3},
		{name: "Bandedmail", cost: 75, damage: 0, armor: 4},
		{name: "Platemail", cost: 102, damage: 0, armor: 5},
	}

	rings := []item{
		{name: "Damage +1", cost: 25, damage: 1, armor: 0},
		{name: "Damage +2", cost: 50, damage: 2, armor: 0},
		{name: "Damage +3", cost: 100, damage: 3, armor: 0},
		{name: "Defense +1", cost: 20, damage: 0, armor: 1},
		{name: "Defense +2", cost: 40, damage: 0, armor: 2},
		{name: "Defense +3", cost: 80, damage: 0, armor: 3},
	}

	return shop{
		weapons: weapons,
		armor:   armor,
		rings:   rings,
	}
}

type fighter struct {
	hitpoints int
	damage    int
	armor     int
}

type shop struct {
	weapons []item
	armor   []item
	rings   []item
}

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type spell struct {
	name   string
	mana   int
	damage int
}
