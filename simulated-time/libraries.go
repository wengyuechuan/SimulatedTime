package simulated_time

import (
	"time"
)

// InitSimulatedTime 初始化模拟时间
func InitSimulatedTime(startTime time.Time, multiple int) (*SimulatedTime, error) {
	if multiple <= 0 {
		ErrMultiple := &ErrorMultiple{multiple: multiple}
		return nil, ErrMultiple
	}
	return &SimulatedTime{
		startTime: startTime,
		multiple:  multiple,
	}, nil
}

// Now 返回当前模拟时间
func (t *SimulatedTime) Now() time.Time {
	if t.elapsedTime == 0 {
		return t.startTime
	}
	return t.startTime.Add(t.elapsedTime)
}

// Add 增加模拟时间
func (t *SimulatedTime) Add(d time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.elapsedTime += d
}

// RunSimulation 模拟时间流逝的协程
func (t *SimulatedTime) RunSimulation() {
	t.mutex.Lock()
	t.running = true
	t.mutex.Unlock()
	for t.running {
		t.Add(time.Second)
		time.Sleep(time.Second / time.Duration(t.multiple))
	}
}

// SetInitialTime 设置模拟时间初始值
func (t *SimulatedTime) SetInitialTime(newStartTime time.Time) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.startTime = newStartTime
	t.elapsedTime = 0
}

// TimePause 暂停模拟时间
func (t *SimulatedTime) TimePause() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.running {
		t.running = false
	}
}
