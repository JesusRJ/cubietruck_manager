package cubietruck

const (
	LED_WHITE  = "/sys/class/leds/cubietruck:white:usr/brightness"
	LED_ORANGE = "/sys/class/leds/cubietruck:orange:usr/brightness"
	LED_GREEN  = "/sys/class/leds/cubietruck:green:usr/brightness"
	LED_BLUE   = "/sys/class/leds/cubietruck:blue:usr/brightness"
	CPU_TEMP   = "/sys/class/thermal/thermal_zone0/temp"
)

// Devices
var (
	Leds = [4]string{LED_BLUE, LED_ORANGE, LED_WHITE, LED_GREEN}
)
