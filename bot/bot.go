package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"

	"github.com/gen2brain/beeep"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

func main() {
	checkAdmin()
}

func checkAdmin() {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		// Not an admin
		notification("Run as admin to get a better analysis of your computer.")
		os.Exit(0)
	}
	// Is an admin
	detectOs()
	sysInfo()
}

func notification(msg string) {
	err := beeep.Alert("Alert", msg, "")

	if err != nil {
		log.Fatal(err)
	}
}

func get_username() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Name)
}

func detectOs() {
	sys := runtime.GOOS

	if sys != "windows" {
		os.Exit(0)
	}
	return
}

func sysInfo() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	// change \\ to / on unix
	diskStat, _ := disk.Usage("\\")

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	fmt.Print(info.Hostname)
}
