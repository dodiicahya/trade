package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cekArr := []int{5, 4, 3, 2, 1}
		diff := Calculate(cekArr)
		exp := &Calculation{
			MaxProfit: 0,
			BuyHour:   1,
			SellHour:  1,
		}
		assert.Equal(t, exp, diff)
	})
	t.Run("#2.", func(t *testing.T) {
		cekArr := []int{3, 2, 1, 5, 6, 2}
		diff := Calculate(cekArr)
		exp := &Calculation{
			MaxProfit: 5,
			BuyHour:   3,
			SellHour:  5,
		}
		assert.Equal(t, exp, diff)
	})
}
