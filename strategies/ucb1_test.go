package strategies

import (
	"github.com/peleteiro/bandit-server/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var strgy = NewUCB1()

func TestReturnTheOneWithZeroHits(t *testing.T) {
	repo := repository.NewMemory()
	repo.Hit("exp", "arm1")

	choosenOne := strgy.Choose(repo, "exp", []string{"arm1", "armWithZeroHits"})

	assert.Equal(t, choosenOne, "armWithZeroHits")
}
