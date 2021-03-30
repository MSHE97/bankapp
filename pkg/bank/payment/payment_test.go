package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 9_99,
		},
		{
			ID:     2,
			Amount: 150_00,
		},
		{
			ID:     3,
			Amount: 100_01,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)

	// Output: 2
}

func ExamplePaymentSources()  {
	payCards := []types.Card{
		{
			PAN: "1xxx xxxx xxxx xxxx",
			Balance: 10_000_00,
			Active: true,
		},
		{
			PAN: "2xxx xxxx xxxx xxxx",
			Balance: 5_000_00,
			Active: true,
		},
		{
			PAN: "3xxx xxxx xxxx xxxx",
			Balance: 20_000_00,
			Active: false,
		},
		{
			PAN: "4xxx xxxx xxxx xxxx",
			Balance: 0,
			Active: true,
		},
	}
	paySources := PaymentSources(payCards)
	
	for _, paySource := range paySources {
		fmt.Println(paySource.Number)	
	}

	// Output: 
	// 1xxx xxxx xxxx xxxx
	// 2xxx xxxx xxxx xxxx
}