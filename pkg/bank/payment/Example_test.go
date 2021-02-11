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
