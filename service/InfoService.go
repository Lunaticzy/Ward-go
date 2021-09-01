package service

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"ward-go/common"
)

var Info = InfoService{}

type InfoService struct {
}

const (
	TiB = 1024 * GiB
	GiB = 1024 * MiB
	MiB = 1024 * 1024
)

// GetInfo info
func (s InfoService) GetInfo() (info common.Info) {

	info.Machine = s.GetMachineInfo()
	info.Uptime = s.GetUptimeInfo()
	info.Processor = s.GetProcessorInfo()
	info.Storage = s.GetStorageInfo()
	info.Setup = common.Setup{ServerName: common.Config.ServerName}
	info.Project = common.Project{Version: common.Config.Version}

	return

}

// GetProcessorInfo cpu
func (s InfoService) GetProcessorInfo() (processor common.Processor) {
	cpuInfo, _ := cpu.Info()
	CpuCounts, _ := cpu.Counts(false)
	processor.CoreCount = fmt.Sprintf(
		"%d Cores",
		CpuCounts,
	)
	processor.ClockSpeed = fmt.Sprintf(
		"%.1fGhz",
		cpuInfo[0].Mhz/1000,
	)
	processor.Name = cpuInfo[0].ModelName
	processor.BitDepth = "64-bit"

	return
}

// GetUptimeInfo uptime
func (s InfoService) GetUptimeInfo() (uptime common.Uptime) {
	hostInfo, _ := host.Info()
	// uptime
	uptimeInSeconds := hostInfo.Uptime
	uptime.Days = fmt.Sprintf("%02d", uptimeInSeconds/86400)
	uptime.Hours = fmt.Sprintf("%02d", (uptimeInSeconds%86400)/3600)
	uptime.Minutes = fmt.Sprintf("%02d", (uptimeInSeconds/60)%60)
	uptime.Seconds = fmt.Sprintf("%02d", uptimeInSeconds%60)

	return
}

func (s InfoService) GetStorageInfo() (storage common.Storage) {

	partitions, _ := disk.Partitions(false)
	SwapInfo, _ := mem.SwapMemory()

	var total, used uint64

	for _, partition := range partitions {
		usage, _ := disk.Usage(partition.Mountpoint)
		total += usage.Total
		used += usage.Used
	}

	storage.DiskUsed = int((used * 100.) / total)

	var storageTotal string

	if total >= TiB {
		storageTotal = fmt.Sprintf("%d GiB Total", total/TiB)
	} else if total >= GiB {
		storageTotal = fmt.Sprintf("%d GiB Total", total/GiB)
	} else {
		storageTotal = fmt.Sprintf("%d GiB Total", total/MiB)
	}
	storage.Total = storageTotal
	storage.DiskCount = fmt.Sprintf("%d Disks", len(partitions))

	var swap string
	if SwapInfo.Total >= GiB {
		swap = fmt.Sprintf("%d GiB Swap", SwapInfo.Total/GiB)
	} else {
		swap = fmt.Sprintf("%d MiB Swap", SwapInfo.Total/MiB)
	}
	storage.SwapAmount = swap
	storage.MainStorage = "Standard disk drives"

	return
}

// GetMachineInfo machine
func (s InfoService) GetMachineInfo() (machine common.Machine) {
	memoryInfo, _ := mem.VirtualMemory()
	hostInfo, _ := host.Info()
	var totalRam string
	if memoryInfo.Total >= GiB {
		totalRam = fmt.Sprintf(
			"%d GiB Ram",
			memoryInfo.Total/GiB,
		)
	} else {
		totalRam = fmt.Sprintf(
			"%d MiB Ram",
			memoryInfo.Total/MiB,
		)
	}

	machine.TotalRAM = totalRam
	machine.OperatingSystem = fmt.Sprintf(
		"%s %s",
		hostInfo.Platform,
		hostInfo.PlatformVersion,
	)
	machine.ProcCount = fmt.Sprintf(
		"%d Procs",
		hostInfo.Procs,
	)
	machine.RAMTypeOrOSBitDepth = "Other"

	return
}
