package money

type Money struct {
	amount   int
	currency string
}

func NewDollar(amount int) Money {
	return Money{
		amount,
		"USD",
	}
}

func NewFranc(amount int) Money {
	return Money{
		amount,
		"CHF",
	}
}

func (d Money) Times(multiplier int) Money {
	return Money{
		d.amount * multiplier,
		d.currency,
	}
}

func (d Money) Equals(money Money) bool {
	return d.amount == money.amount && d.currency == money.currency
}

func (d Money) Currency() string {
	return d.currency
}
