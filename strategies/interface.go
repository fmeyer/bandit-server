package strategies

import "github.com/peleteiro/bandit-server/repository"

type Strategy interface {
	Choose(repo repository.Repository, context string, experiments []string) string
}
