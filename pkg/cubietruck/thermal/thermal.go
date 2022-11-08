package thermal

import "github.com/jesusrj/cubietruck/pkg/cubietruck/internal/device/thermal"

const (
	cpu_temp = "/sys/class/thermal/thermal_zone0/temp"
)

var (
	cpuTemp = thermal.New(cpu_temp)
)
