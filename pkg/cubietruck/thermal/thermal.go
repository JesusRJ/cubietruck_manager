package thermal

import (
	"strconv"
	"strings"

	"github.com/jesusrj/cubietruck/pkg/cubietruck/internal/device/thermal"
)

const (
	cpu_temp = "/sys/class/thermal/thermal_zone0/temp"
)

var (
	cpuTemp = thermal.New(cpu_temp)
)

func CPUTemp() (uint, error) {
	b, err := cpuTemp.Read()
	if err != nil {
		return 0, err
	}
	// trim carriage return
	x := strings.TrimSuffix(string(b), "\n")
	temp, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}
	return uint(temp), nil
}
