package money

type Expression interface {
	Reduce(b Bank, to string) Money
}

type Sum struct {
	Augend, Addend Expression
}

func (s Sum) Reduce(b Bank, to string) Money {
	return Money{
		s.Augend.Reduce(b, to).amount + s.Addend.Reduce(b, to).amount,
		to,
	}
}
