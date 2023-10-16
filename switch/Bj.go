package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch  {
	case card =="ace":
		return 11
	case card == "jack" || card =="queen" || card == "king":
		return 10
	case  card == "two":
		return 2
	case  card =="three":
		return 3
	case  card =="four":
		return 4
	case  card == "five":
		return 5
	case  card == "six":
		return 6
	case card == "seven":
		return 7
	case  card == "eight":
		return 8
	case  card == "nine":
		return 9
	case  card == "ten":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	score := ParseCard(card1) + ParseCard(card2)
	switch {
	case score == 22:
		return "P"
	case score == 21:
		if ParseCard(dealerCard) != 10 && ParseCard(dealerCard) != 11 {
			return "W"
		} else {
			return "S"
		}
	case score >= 12 && score <= 16:
		if 7 <= ParseCard(dealerCard) {
			return "H"
		} else {
			return "S"
		}
	case score >= 17 && score <= 20:
		return "S"
	case score <= 11:
		return "H"
	default:
		return "H"
	}

}
