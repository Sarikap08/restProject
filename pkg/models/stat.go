package models

type Stat struct {
	CpuTemp  int `json:"cpuTemp"`
	FanSpeed int `json:"fanSpeed"`
	HDDSpace int `json:"HDDSpace"`
}
