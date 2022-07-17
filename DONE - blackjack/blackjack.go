package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) (val int) {
	switch card {
	case "ace":
		val = 11

	case "two":
		val = 2

	case "three":
		val = 3

	case "four":
		val = 4

	case "five":
		val = 5

	case "six":
		val = 6

	case "seven":
		val = 7

	case "eight":
		val = 8

	case "nine":
		val = 9

	case "ten", "queen", "jack", "king":
		val = 10

	default:
		val = 0
	}
	return
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) (action string) {
	if card1 == "ace" && card2 == "ace" {
		action = "P"
	} else {
		c1, c2, dc := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
		// cards := []int{c1, c2, dc}
		switch {
		case (c1+c2 == 21):
			if dc != 11 && dc != 10 {
				action = "W"
			} else {
				action = "S"
			}
		case c1+c2 >= 17 && c1+c2 <= 20:
			action = "S"
		case c1+c2 >= 12 && c1+c2 <= 16:
			if dc >= 7 {
				action = "H"
			} else {
				action = "S"
			}
		case c1+c2 <= 11:
			action = "H"
		}

	}

	return
}
