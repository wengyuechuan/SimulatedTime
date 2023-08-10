package main

import (
	"fmt"
	SimulatedTime "github.com/wengyuechuan/SimulatedTime/simulated-time"
	"time"
)

func main() {
	simulatedTime, err := SimulatedTime.InitSimulatedTime(time.Now(), 20)
	if err != nil {
		panic(err)
	}
	if simulatedTime == nil {
		panic("模拟时间为空")
	}
	// 打印时间
	fmt.Println(simulatedTime.Now())
	// 模拟时间流逝
	go simulatedTime.RunSimulation()
	// 命令行工具
	go SimulatedTime.RunSimulationTool(simulatedTime)

	time.Sleep(time.Second * 100)
}
