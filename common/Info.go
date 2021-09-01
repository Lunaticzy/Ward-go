package common

type Info struct {
	Processor Processor `json:"processor"`
	Machine   Machine   `json:"machine"`
	Storage   Storage   `json:"storage"`
	Uptime    Uptime    `json:"uptime"`
	Setup     Setup     `json:"setup"`
	Project   Project   `json:"project"`
}
type Processor struct {
	Name       string `json:"name"`
	CoreCount  string `json:"coreCount"`
	ClockSpeed string `json:"clockSpeed"`
	BitDepth   string `json:"bitDepth"`
}
type Machine struct {
	OperatingSystem     string `json:"operatingSystem"`
	TotalRAM            string `json:"totalRam"`
	RAMTypeOrOSBitDepth string `json:"ramTypeOrOSBitDepth"`
	ProcCount           string `json:"procCount"`
}
type Storage struct {
	MainStorage string `json:"mainStorage"`
	Total       string `json:"total"`
	DiskCount   string `json:"diskCount"`
	DiskUsed    int    `json:"diskUsed"`
	SwapAmount  string `json:"swapAmount"`
}
type Uptime struct {
	Days    string `json:"days"`
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
	Seconds string `json:"seconds"`
}
type Setup struct {
	ServerName string `json:"serverName"`
}
type Project struct {
	Version string `json:"version"`
}
