package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/33r01b/systray/modules"
)

var (
	cpu       string
	temp      string
	mem       string
	diskspace string
	network   string
	bat       string
	date      string
)

func init() {
	cpu = "cpu"
	temp = "temp"
	mem = "mem"
	diskspace = "diskspace"
	network = "network"
	bat = "bat"
	date = "date"
}

func main() {
	go cpuUpdate()
	go tempUpdate()
	go memUpdate()
	go diskUpdate()
	go networkUpdate()
	go batUpdate()
	go dateUpdate()
	// TODO bluetooth

	var stateBuilder strings.Builder

	for {
		stateBuilder.Reset()

		stateBuilder.WriteString(cpu)
		stateBuilder.WriteString("  ")
		stateBuilder.WriteString(temp)
		stateBuilder.WriteString("  ")
		stateBuilder.WriteString(mem)
		stateBuilder.WriteString("  ")
		stateBuilder.WriteString(diskspace)
		stateBuilder.WriteString("  ")
		stateBuilder.WriteString(network)
		stateBuilder.WriteString("  ")
		stateBuilder.WriteString(bat)
		stateBuilder.WriteString("  ")
		stateBuilder.WriteString(date)

		fmt.Println(stateBuilder.String())

		time.Sleep(1 * time.Second)
	}
}

func cpuUpdate() {
	for {
		value, err := modules.Cpu()
		if err != nil {
			log.Println(err)
			return
		}

		cpu = fmt.Sprintf("cpu %d%%", value)

		time.Sleep(5 * time.Second)
	}
}

func tempUpdate() {
	for {
		value, err := modules.Temp()
		if err != nil {
			log.Println(err)
			return
		}

		temp = fmt.Sprintf("%sC", string(value))

		time.Sleep(3 * time.Second)
	}
}

func memUpdate() {
	for {
		value, err := modules.Mem()
		if err != nil {
			log.Println(err)
			return
		}

		mem = fmt.Sprintf("%d MiB", value)

		time.Sleep(5 * time.Second)
	}
}

func diskUpdate() {
	for {
		available, err := modules.DiskAvailable()
		if err != nil {
			log.Println(err)
			return
		}

		diskspace = fmt.Sprintf("%dG", available)

		time.Sleep(5 * time.Minute)
	}
}

func networkUpdate() {
	for {
		down, up, err := modules.Network()
		if err != nil {
			log.Println(err)
			return
		}

		network = fmt.Sprintf("down %d kB/s up %d kB/s", down, up)

		time.Sleep(1 * time.Second)
	}
}

func batUpdate() {
	for {
		capacity, err := modules.BatCapacity()
		if err != nil {
			log.Println(err)
			return
		}

		isCharging, err := modules.BatCharging()
		if err != nil {
			log.Println(err)
			return
		}

		charge := ""
		if isCharging {
			charge = "+"
		}

		bat = fmt.Sprintf("bat%s %s", charge, capacity)

		time.Sleep(10 * time.Second)
	}
}

func dateUpdate() {
	for {
		date = time.Now().Format("Mon 01-02 15:04")

		time.Sleep(10 * time.Second)
	}
}
