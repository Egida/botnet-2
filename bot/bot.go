package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/tatsushid/go-fastping"
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
	ping("1.1.1.1", 1)
}

func notification(msg string) {
	err := beeep.Alert("Alert", msg, "")

	if err != nil {
		log.Fatal(err)
	}
}

func ping(ip string, pingCount int) {
	for i := 0; i < pingCount; i++ {
		p := fastping.NewPinger()
		ra, err := net.ResolveIPAddr("ip4:icmp", ip)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			fmt.Printf("\n[+] IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		}
		p.OnIdle = func() {
			fmt.Println("\n[+] Ping completed!")
		}
		err = p.Run()
		if err != nil {
			fmt.Println(err)
		}
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
