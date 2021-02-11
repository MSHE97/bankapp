package main

import (
	// "bank/pkg/bank/deposit"
	// "bank/pkg/bank/transfer"
	"bank/pkg/bank/types"
	"bank/pkg/bank/card"
	"fmt"
)

func main() {
	// min, max := deposit.Calculate(500_000_00, "TJS")
	// fmt.Println(min)
	// fmt.Println(max)
	// transferMoney := 1_025_00
	// fmt.Println(transfer.Total(transferMoney))
	card0 := types.Card {
		ID:			12345,
		PAN: 		"1234 xxxx xxxx 5678",
		Balance: 	999_99,
		Currency: 	"TJS",
		Color:		"White",
		Name:		"Infinity",
		Active:		true,
	}
	card := card.IssueCard(types.USD, "Platinum", "Universe")
	fmt.Printf("%+v", card)
	card0 = handle(card0)
	fmt.Printf("\n%+v", card0)
}

func handle(card types.Card) types.Card {
	card.Active = false
	return card // Todo
}