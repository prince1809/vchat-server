package model

import "time"

type ScheduledTask struct {
	Name      string        `json:"name"`
	Interval  time.Duration `json:"interval"`
	Recurring bool
	function  func()
	cancel    chan struct{}
	cancelled chan struct{}
}
