package connection

import (
	"fmt"
	"strconv"
)

var (
	mega int64 = 1000 * 1000
	giga int64 = 1000 * mega
)

func speedIntToStr(speed int64) (string, error) {
	var base int64
	var unit string

	if (speed % giga) == 0 {
		unit = "Gbps"
		base = speed / giga
	} else if (speed % mega) == 0 {
		unit = "Mbps"
		base = speed / mega
	} else {
		return "", fmt.Errorf("unsupported speed value %v", speed)
	}
	return strconv.Itoa(int(base)) + unit, nil
}
