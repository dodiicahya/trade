package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMatchFactor(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cur, _ := FindFactorSix("12")
		assert.Equal(t, 1, cur)
	})
}
