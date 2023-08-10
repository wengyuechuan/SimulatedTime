package test

import (
	SimulatedTime "github.com/wengyuechuan/SimulatedTime/simulated-time"
	"testing"
	"time"
)

// TestInitSimulatedTime 测试初始化
func TestInitSimulatedTime(t *testing.T) {
	t.Log("测试初始化")
	simulatedTime, err := SimulatedTime.InitSimulatedTime(time.Now(), 20)
	if err != nil {
		t.Error(err)
	}
	if simulatedTime == nil {
		t.Error("模拟时间为空")
	}
	// 打印时间
	t.Log(simulatedTime.Now())
}
