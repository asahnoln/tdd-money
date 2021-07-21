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

func (d Money) Equals(m Money) bool {
	return d.amount == m.amount && d.currency == m.currency
}

func (d Money) Currency() string {
	return d.currency
}

func (d Money) Plus(m Money) Sum {
	return Sum{d, m}
}

func (d Money) Reduce(to string) Money {
	return d
}
