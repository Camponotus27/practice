package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	for {
		cpuUsage, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Printf("Error al obtener el uso de la CPU: %v\n", err)
			continue
		}

		memoryUsage, err := mem.VirtualMemory()
		if err != nil {
			fmt.Printf("Error al obtener el uso de la memoria: %v\n", err)
			continue
		}

		if len(cpuUsage) > 0 {
			fmt.Printf("Uso de la CPU: %.2f%%\n", cpuUsage[0])
			fmt.Printf("Memoria total: %v, Memoria usada: %v, Porcentaje de memoria usada: %.2f%%\n", memoryUsage.Total, memoryUsage.Used, memoryUsage.UsedPercent)
		}

		time.Sleep(time.Second * 5)
	}
}
