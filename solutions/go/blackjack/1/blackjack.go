package blackjack

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "joker":
		return 0
	default:
		panic("Invalid card: " + card)
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case playerScore == 21 && dealerScore < 10:
		return "W"
	case playerScore == 21:
		return "S"
	case playerScore >= 17:
		return "S"
	case playerScore >= 12 && dealerScore >= 7:
		return "H"
	case playerScore >= 12:
		return "S"
	default:
		return "H"
	}
}
