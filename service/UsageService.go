package service

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
	"ward-go/common"
)

var Usage = UsageService{}

type UsageService struct{}

func (u UsageService) GetUsage() (usage common.Usage) {

	processorPercent, _ := cpu.Percent(time.Millisecond*500, false)
	usage.Processor = int(processorPercent[0])

	memInfo, _ := mem.VirtualMemory()
	usage.Ram = int(memInfo.UsedPercent)

	usage.Storage = Info.GetStorageInfo().DiskUsed

	return
}
