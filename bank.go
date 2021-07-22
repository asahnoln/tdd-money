package money

type Bank struct {
	rates map[string]int
}

func NewBank() Bank {
	return Bank{map[string]int{}}
}

func (b Bank) Reduce(e Expression, to string) Money {
	return e.Reduce(b, to)
}

func (b Bank) AddRate(from, to string, rate int) {
	b.rates[from + to] = rate
}

func (b Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}

	rate := b.rates[from + to]
	return rate
}
