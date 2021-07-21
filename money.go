package currency

type Money struct{
	amount int
	name string
}

func NewDollar(amount int) Money {
	return Money{
		amount: amount,
		name:   "dollar",
	}
}

func NewFranc(amount int) Money {
	return Money{
		amount: amount,
		name:   "franc",
	}
}

func (d Money) Times(multiplier int) Money {
	return Money{
		amount: d.amount * multiplier,
		name:   d.name,
	}
}

func (d Money) Equals(money Money) bool {
	return d.amount == money.amount && d.name == money.name
}
