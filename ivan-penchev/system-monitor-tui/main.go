package main

import (
	"fmt"
)

func main() {
	printSystemInfo()
}

func printSystemInfo() {
	cpuUsage, _ := GetCPUStats()
	memoryUsage, _ := GetMEMStats()
	runningProcesses, _ := GetProcesses(10) // Get top 10 CPU intensive processes

	fmt.Println("CPU Percentage    :", cpuUsage)
	fmt.Println("Memory Percentage :", memoryUsage)
	fmt.Println("Running Processes :", runningProcesses)
}

