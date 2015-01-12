package repository

type ExperimentData struct {
	TotalHits int64
	Arms      map[string]ArmData
}

type ArmData struct {
	Hits    int64
	Rewards int64
}

type Repository interface {
	Get(experiment string, arms []string) ExperimentData
	Hit(experiment string, arm string)
	Reward(experiment string, arm string)
}
