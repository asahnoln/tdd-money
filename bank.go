package money

type Bank struct {
}

func (b Bank) Reduce(e Expression, to string) Money {
	return e.Reduce(to)
}
