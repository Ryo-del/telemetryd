package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	x, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x[0])
}
