package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	for quantity_repetitions := 1; quantity_repetitions <= 3; quantity_repetitions++ {
		var y bool = false
		x, err := cpu.Percent(time.Second, y)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(x[0])

		time.Sleep(3 * time.Second)
	}
	fmt.Println("end")
}
