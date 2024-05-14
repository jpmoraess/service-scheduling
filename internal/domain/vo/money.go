package vo

import (
	"fmt"
	"math/big"
)

var ZERO = NewMoney(0)

type Money struct {
	amount *big.Float
}

func NewMoney(amount float64) *Money {
	return &Money{amount: big.NewFloat(amount).SetPrec(64)}
}

func (m *Money) IsGreaterThan(other *Money) bool {
	return m.amount.Cmp(other.amount) > 0
}

func (m *Money) Add(other *Money) *Money {
	result := new(big.Float).Add(m.amount, other.amount)
	return &Money{amount: result}
}

func (m *Money) Subtract(other *Money) *Money {
	result := new(big.Float).Sub(m.amount, other.amount)
	return &Money{amount: result}
}

func (m *Money) Multiply(multiplier int64) *Money {
	multiplierBig := big.NewFloat(float64(multiplier)).SetPrec(64)
	result := new(big.Float).Mul(m.amount, multiplierBig)
	return &Money{amount: result}
}

func (m *Money) GetAmount() *big.Float {
	return m.amount
}

func (m *Money) GetAmountFloat64() float64 {
	f, _ := m.amount.Float64()
	return f
}

func (m *Money) Equals(other *Money) bool {
	return m.amount.Cmp(other.amount) == 0
}

func (m *Money) String() string {
	return fmt.Sprintf("%.2f", m.amount)
}
