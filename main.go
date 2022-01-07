package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
    cpu string
    temp string
    mem string
    diskspace string
    network string
    wifi string
    bat string
    date string
    bluetooth string
)

func init()  {
    cpu = "cpu"
    temp = "temp"
    mem = "mem"
    diskspace = "diskspace"
    network = "network"
    wifi = "wifi"
    bat = "bat"
    date = "date"
    bluetooth = "bluetooth"
}

func main() {

    cpuUpdate()

    var stateBuilder strings.Builder
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
    stateBuilder.WriteString(wifi)
    stateBuilder.WriteString("  ")
    stateBuilder.WriteString(bat)
    stateBuilder.WriteString("  ")
    stateBuilder.WriteString(date)
    stateBuilder.WriteString("  ")
    stateBuilder.WriteString(bluetooth)

    fmt.Println(stateBuilder.String())
}

func cpuUpdate () {
    prevTotal, prevIdle, err := getCpuStat()
    if err != nil {
        log.Fatal(err)
    }

    time.Sleep(500 * time.Millisecond)

    total, idle, err := getCpuStat()
    if err != nil {
        log.Fatal(err)
    }

    var cpuUsage int
    if (total-prevTotal) > 0 {
        cpuUsage = 100 * ((total - prevTotal) - (idle-prevIdle)) / (total - prevTotal) 
    }

    cpu = fmt.Sprintf("cpu %d", cpuUsage)
}

func getCpuStat() (total, idle int, err error) {
    stat, err := os.Open("/proc/stat")
    if err != nil {
        return 0, 0, fmt.Errorf("cannot open /proc/stat: %w", err)
    }
    defer stat.Close()

    reader := bufio.NewReader(stat)
    ln, _, err := reader.ReadLine()
    if err != nil {
        return 0, 0, fmt.Errorf("cannot read /proc/stat: %w", err)
    }

    cols := strings.Fields(string(ln))
    if len(cols) < 5 {
        return 0, 0, errors.New("stat fields length mismatch")
    }

    userProcesses, err := strconv.Atoi(cols[1])
    if err != nil {
        return 0, 0, fmt.Errorf("cannot parse cpu 'userProcesses' state: %w", err)
    }

    nicedProcesses, err := strconv.Atoi(cols[2])
    if err != nil {
        return 0, 0, fmt.Errorf("cannot parse cpu 'nicedProcesses' state: %w", err)
    }

    systemProcesses, err := strconv.Atoi(cols[3])
    if err != nil {
        return 0, 0, fmt.Errorf("cannot parse cpu 'systemProcesses' state: %w", err)
    }

    idle, err = strconv.Atoi(cols[4])
    if err != nil {
        return 0, 0, fmt.Errorf("cannot parse cpu 'idle' state: %w", err)
    }

    total = userProcesses + nicedProcesses + systemProcesses + idle
    
    return total, idle, err
}
