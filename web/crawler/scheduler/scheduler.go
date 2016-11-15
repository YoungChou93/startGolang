package scheduler

import "container/list"
type Scheduler struct {
	queue *list.List
}

func NewScheduler() *Scheduler {
	queue := list.New()
	return &Scheduler{queue: queue}
}
