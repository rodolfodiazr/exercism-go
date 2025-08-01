package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	numbers := map[string]int{
		"two": 2, "three": 3, "four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9, "ten": 10,
		"jack": 10, "queen": 10, "king": 10, "ace": 11,
	}

	if value, exists := numbers[card]; exists {
		return value
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var decision string
    switch {
        case card1 == "ace" && card2 == "ace": 
        	decision = "P"
        	break;
        case ((ParseCard(card1) + ParseCard(card2)) == 21 && ParseCard(dealerCard) < 10):
        	decision = "W"
        	break;
        case ((ParseCard(card1) + ParseCard(card2)) == 21 && ParseCard(dealerCard) >= 10):
        	decision = "S"
        	break;
        case ((ParseCard(card1) + ParseCard(card2)) >= 17 && (ParseCard(card1) + ParseCard(card2)) <= 20):
        	decision = "S"
        	break;
        case ((ParseCard(card1) + ParseCard(card2)) >= 12 && (ParseCard(card1) + ParseCard(card2)) <= 16 && ParseCard(dealerCard) < 7):
        	decision = "S"
        	break;
        case ((ParseCard(card1) + ParseCard(card2)) >= 12 && (ParseCard(card1) + ParseCard(card2)) <= 16 && ParseCard(dealerCard) >= 7):
        	decision = "H"
        	break;
        case ((ParseCard(card1) + ParseCard(card2)) <= 11):
        	decision = "H"
        	break;
    }
	return decision
}
