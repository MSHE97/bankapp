package card

import (
	"bank/pkg/bank/types"
)

const RubCource = 0.1510
const UsdCource = 11.3200

//Total вычисляет сумму денег на всех активных картах с положительным балансом
func Total(cards []types.Card)  int64 {
	sum := int64(0)

	for _, card := range cards {
		if checkOut(card, card.Balance, "") > 0 {
			sum += int64(card.Balance)
		}
	}
	return sum
}

// IssueCard Возвращает заполненную структуру Card с заданной валютой, цветом и именем карты 
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card {
	ID:	      10000,
	PAN:      "5058 xxxx xxxx 0001",
	Balance:  0,
	Currency: currency,
	Color:    color,
	Name:     name,
	Active:   true,
	MinMoney: 0,
	}
	return card
}

// Withdraw Снимает деньги с активной карты ( < 20_000 in TJS )
func Withdraw(card *types.Card, amount types.Money){
	if checkOut(*card, amount, "withdraw") > 0 {
		card.Balance -= amount
	}
}

// Deposit зачисляет деньги на активную карту ( < 50_000 in TJS)
func Deposit(card *types.Card, amount types.Money) {
	if checkOut(*card, amount, "deposit") > 0 {
		card.Balance += amount
	}
}

// AddBonus зачисляет бонусные проценты на минимальный остаток счёта (бонус не более 50_000_00 TJS)
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	amount := ( ( card.MinMoney * types.Money(percent * daysInMonth) ) / types.Money(daysInYear * 100) )
	
	err := checkOut( *card, amount, "bonus" )
	if err == -3  {
		switch card.Currency {
		case types.RUB:
			rub := 5_000_00 / RubCource
			amount = types.Money(rub)
		case types.USD:
			dol := 5_000_00 / UsdCource
			amount = types.Money(dol)
		default:
			amount = 5_000_00
		}
		err += 4
	}
	
	if err > 0 {
		card.Balance += amount
	}
}

// checkOut проверит состояние карты; даёт добро на снятие/зачисление денег, если не обнаружит проблем
func checkOut( card types.Card, amount types.Money, operation string ) int {

	if !card.Active {
		return -2
	}

	if amount <= 0 {
		return -1
	}

	switch operation {
	case "withdraw":
		
		if card.Balance - amount < 0 {
			return 0
		}

		if inTJS(card.Currency, amount) > 20_000_00 {
			return -3
		}
	case "deposit":

		if inTJS( card.Currency, amount ) > 50_000_00 {
			return -3
		}
	case "bonus":

		if card.Balance <= 0 {
			return -4
		}
		if inTJS( card.Currency, amount ) > 5_000_00 {
			return -3
		}
	}
	return 1
}

// inTJS приведёт снимаемые деньги в TJS
func inTJS(currency types.Currency ,someMoney types.Money) types.Money{
	switch currency {
	case types.RUB:
		return types.Money(float64(someMoney) * RubCource)
	case types.USD:
		return types.Money(float64(someMoney) * UsdCource)
	}
	return someMoney
}