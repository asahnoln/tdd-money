package money

type Expression interface {
	Reduce(to string) Money
}

type Sum struct {
	Augend, Addend Money
}

func (s Sum) Reduce(to string) Money {
	return Money{
		s.Augend.amount + s.Addend.amount,
		to,
	}
}
