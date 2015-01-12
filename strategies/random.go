package strategies

import (
	"github.com/peleteiro/bandit-server/repository"
	"math/rand"
)

type Random struct{}

func NewRandom() Random {
	return Random{}
}

func (_ Random) Choose(repo repository.Repository, experiment string, arms []string) string {
	var arm = getRandomArm(arms)
	repo.Hit(experiment, arm)
	return arm
}

func getRandomArm(arms []string) string {
	var i = randInt(len(arms))
	return arms[i]
}

func randInt(len int) int {
	return rand.Intn(len)
}
