package types

// Money - денежная суммы в минимальных еденицах (дирамы, копейки, центы и т.д.)
type Money int64

// Currency - код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN номер карты
type PAN string

// Card - содержит информацию о платёжной карте
type Card struct {
	ID 			int
	PAN
	Balance 	Money
	Currency
	Color 		string
	Name 		string
	Active 		bool
	MinMoney	Money
}

type Payment struct {
	ID int
	Amount Money
}

type PaymentSource struct {
	Type string // 'card'
    Number string // номер вида '5058 xxxx xxxx 8888' 
	Balance Money // баланс в дирамах
}