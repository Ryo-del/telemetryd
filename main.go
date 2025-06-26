package main

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	for quantity_repetitions := 1; quantity_repetitions <= 30; quantity_repetitions++ {
		//%cpu
		cpuloads, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println(err)
		}
		//%ram
		ramloads, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println(err)
		}
		cpuload := cpuloads[0]
		cpuload = math.Round(cpuload)
		ramload := ramloads.UsedPercent
		ramload = math.Round(ramload)

		fmt.Println(ramload)
		fmt.Println(cpuload)

		//tick rate
		time.Sleep(3 * time.Second)
	}
	fmt.Println("end")
}
