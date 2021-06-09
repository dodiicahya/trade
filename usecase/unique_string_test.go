package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueString(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cekArr := "sebaerb"
		cur := FirstOccurence(cekArr)
		assert.Equal(t, "sebar", cur)
	})
	t.Run("#2.", func(t *testing.T) {
		cekArr := "sebaerb"
		cur := LexicoGraphically(cekArr)
		assert.Equal(t, "abers", cur)
	})
}
