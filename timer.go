package gotils

import (
	"time"
)

type Timer struct {
	start time.Time
	Name  string
}

func NewTimer(name string) Timer {
	t := Timer{}
	t.Name = name
	t.start = time.Now()
	return t
}

func (t *Timer) GetDuration() time.Duration {
	n := time.Now()
	return n.Sub(t.start)
}

func (t *Timer) GetDurationString() string {
	return t.GetDuration().String()
}
