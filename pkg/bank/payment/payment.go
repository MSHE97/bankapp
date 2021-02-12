package payment

import "bank/pkg/bank/types"

// Max возвращает максимальный платёж
func Max(payments []types.Payment) types.Payment {
	max := payments[0]

	for _, payment := range payments {
		if max.Amount < payment.Amount {
			max = payment
		}
	}
	return max
}

func PaymentSources(cards []types.Card) []types.PaymentSource {

	paySources := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active == true && card.Balance > 0 {
			paySources = append(paySources, types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			})
		}
	}
	return paySources
}