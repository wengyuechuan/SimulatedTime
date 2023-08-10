package test

import (
	SimulatedTime "github.com/wengyuechuan/SimulatedTime/simulated-time"
	"testing"
	"time"
)

func TestSimulation(t *testing.T) {
	t.Log("测试时间流逝")
	simulatedTime, err := SimulatedTime.InitSimulatedTime(time.Now(), 20)
	if err != nil {
		t.Error(err)
	}
	if simulatedTime == nil {
		t.Error("模拟时间为空")
	}
	// 打印时间
	t.Log(simulatedTime.Now())

	go simulatedTime.RunSimulation()

	// 等待协程运行
	time.Sleep(time.Second * 10)
	t.Log(simulatedTime.Now())
}
