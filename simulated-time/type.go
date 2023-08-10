package simulated_time

import (
	"sync"
	"time"
)

type SimulatedTime struct {
	mutex       sync.Mutex // 加锁，防止多个协程同时访问
	elapsedTime time.Duration
	startTime   time.Time
	multiple    int
	running     bool
}
