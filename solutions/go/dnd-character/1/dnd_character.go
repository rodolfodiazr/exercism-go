package dndcharacter

import (
    "math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
    diff := score - 10
	if diff >= 0 {
		return diff / 2
	}
	return (diff - 1) / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    sum := 0
	min := 7
	for i := 1; i <= 4; i++ {
		randomNumber := rand.Intn(6) + 1
        if randomNumber < min {
			min = randomNumber
		}
		
        sum += randomNumber
	}
	return sum - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
