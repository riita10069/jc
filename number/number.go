package number

import "errors"

type Number struct {
	prime int
	number int
}

type IElement interface {
	Add()
	Sub()
	Mul()
	Div()
}

func (f *Number) Add(other *Number) int {
	return (f.number + other.number)%f.prime
}

func (f *Number) Sub(other *Number) int {
	return (f.number - other.number)%f.prime
}

func (f *Number) Mul(other *Number) int {
	return (f.number * other.number)%f.prime
}

func (f *Number) Div(other *Number) (int, error) {
	if other.number == 0 {
		return -1, errors.New("(*>△<)<0 で割っちゃダメだよ！！")
	}
	for i := range make([]int, f.prime, f.prime) {
		if other.number * i == f.number {
			return i, nil
		}
	}
	return 0, nil
}
