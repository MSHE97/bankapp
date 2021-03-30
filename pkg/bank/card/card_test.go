package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card {
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 5_000_00,
			Active: true,
		},
		{
			Balance: 20_000_00,
			Active: false,
		},
		{
			Balance: -500_00,
			Active: true,
		},
	}

	fmt.Println( Total(cards) )

	//Output: 1500000
}

func ExampleWithdraw_positive() {
	card := types.Card{Currency: types.USD, Balance: 2_00_00, Active: true}
	Withdraw(&card, 2_00_00)
	fmt.Println(card.Balance)

	// Output: 0
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 20_000_00)
	fmt.Println(card.Balance)

	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 2_000_000_00, Active: false}
	Withdraw(&card, 20_000_00)
	fmt.Println(card.Balance)

	// Output: 200000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 2_000_000_00, Active: true}
	Withdraw(&card, 20_000_01)
	fmt.Println(card.Balance)

	// Output: 200000000
}

func ExampleDeposit_positive() {
	card := types.Card{
		Currency: types.TJS,
		Balance:  -50_000_00,
		Active:   true,
	}
	Deposit(&card, 50_000_00)
	fmt.Println(card.Balance)

	// Output: 0
}

func ExampleDeposit_inactive() {
	card := types.Card{
		Currency: types.TJS,
		Balance:  0,
		Active:   false,
	}
	Deposit(&card, 50_000_00)
	fmt.Println(card.Balance)

	// Output: 0
}

func ExampleDeposit_limit() {
	card := types.Card{
		Currency: types.TJS,
		Balance:  0,
		Active:   true,
	}
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)

	// Output: 0
}

func ExampleAddBonus_positive()  {
	card := types.Card{
		Currency: types.TJS,
		Balance:  10_000_00,
		Active:   true,
		MinMoney: 10_000_00,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 1002465
}

func ExampleAddBonus_inactive()  {
	card := types.Card{
		Currency: types.TJS,
		Balance:  10_000_00,
		Active:   false,
		MinMoney: 10_000_00,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 1000000
}

func ExampleAddBonus_balance()  {
	card := types.Card{
		Currency: types.TJS,
		Balance:  -12_00,
		Active:   true,
		MinMoney: -50_00,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: -1200
}

func ExampleAddBonus_inlimit()  {
	card := types.Card{
		Currency: types.TJS,
		Balance:  2_500_000_00,
		Active:   true,
		MinMoney: 2_500_000_00,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 250500000
}