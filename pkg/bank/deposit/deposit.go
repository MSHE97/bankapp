package deposit

import "bank/pkg/bank/types"

// Calculate расчитывает рамки ежемесячного дохода от взноса, в зависимости от валюты
func Calculate(amount types.Money, currency types.Currency)  (min, max types.Money){
	minPersent,	maxPersent := percentFor(currency)
	min = amount * types.Money(minPersent) / (12 * 100)
	max = amount * types.Money(maxPersent) / (12 * 100)
	return
}
// Returns deposit percents for each currency
func percentFor(currency types.Currency) (min int, max int){
	switch currency {
	case types.RUB:
		return 4, 5
	case types.USD:
		return 3, 4
	default:
		return 4, 6 //todo: ?
	}
}