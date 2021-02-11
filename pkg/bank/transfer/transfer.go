package transfer

//import "bank/pkg/bank/types"

const bonusPersent = 0.0050
type Money int64

//Расчитывает по amount начисляемый при переводе бонус
func Bonus(amount Money) Money {
	bonus := Money(float64(amount) * bonusPersent)
	return bonus
}

// Returns total transfered money including 0.5% bonus
func Total(someMoney Money) Money {
	total := someMoney + Bonus(someMoney)
	return total
}