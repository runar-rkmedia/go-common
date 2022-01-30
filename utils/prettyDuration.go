package utils

import (
	"fmt"
	"math"
	"time"
)

// Returns a more humanly readable
func PrettyDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%03dms", d.Milliseconds())
	}
	if d < time.Minute {
		return fmt.Sprintf("%02.1fs", d.Seconds())
	}
	if d < time.Hour {
		_s := d.Seconds()
		m := math.Floor(_s / 60)
		s := _s - (m * 60)
		return fmt.Sprintf("%02.0f:%02.0f", m, s)
	}
	return d.Round(time.Second).String()

}
