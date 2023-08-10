package simulated_time

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// printCommandMeaning 显示指令的含义
func printCommandMeaning(command string, meaning string) {
	fmt.Printf(" %s - %s\n", command, meaning)
}

// printCommandUsage 显示指令的格式和用法
func printCommandUsage(command string, usage string) {
	fmt.Printf(" %s <参数> - %s\n", command, usage)
}

// RunSimulationTool 命令行工具函数
func RunSimulationTool(simTime *SimulatedTime) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("模拟时间工具命令行")
	fmt.Println("-----------------------")
	fmt.Println("命令列表:")
	printCommandMeaning("start", "开始模拟时间")
	printCommandMeaning("stop", "停止模拟时间")
	printCommandMeaning("reset", "重置模拟时间")
	printCommandUsage("set", "设置模拟时间初始值 (格式: 2006-01-02 15:04:05)")
	printCommandMeaning("now", "显示当前模拟时间")
	printCommandMeaning("exit", "退出程序")
	printCommandMeaning("<command>-help", "显示指令的含义和用法")
	fmt.Println("-----------------------")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.HasSuffix(input, "-help") {
			cmd := strings.TrimSuffix(input, "-help")
			switch cmd {
			case "start":
				printCommandMeaning("start", "开始模拟时间")
			case "stop":
				printCommandMeaning("stop", "停止模拟时间")
			case "reset":
				printCommandMeaning("reset", "重置模拟时间")
			case "set":
				printCommandUsage("set", "设置模拟时间初始值 (格式: 2006-01-02 15:04:05)")
			case "now":
				printCommandMeaning("now", "显示当前模拟时间")
			case "exit":
				printCommandMeaning("exit", "退出程序")
			default:
				fmt.Println("无效的命令")
			}
			continue
		}

		switch input {
		case "start":
			if !simTime.running {
				simTime.running = true
				fmt.Println("模拟时间已开始")
			} else {
				fmt.Println("模拟时间已经在运行中")
			}
		case "stop":
			simTime.TimePause()
		case "reset":
			simTime.SetInitialTime(time.Now())
			fmt.Println("模拟时间已重置")
		case "now":
			fmt.Printf("当前模拟时间：%v\n", simTime.Now())
		case "exit":
			simTime.running = false
			fmt.Println("程序已退出")
			return
		default:
			if strings.HasPrefix(input, "set") {
				parts := strings.Split(input, " ")
				if len(parts) == 3 {
					timeStr := strings.TrimSpace(parts[1] + " " + parts[2])
					newStartTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
					if err == nil {
						simTime.SetInitialTime(newStartTime)
						fmt.Println("模拟时间初始值已设置")
					} else {
						fmt.Println("无效的时间格式")
					}
				} else {
					fmt.Println("无效的命令格式")
				}
			} else {
				fmt.Println("无效的命令")
			}
		}
	}
}
