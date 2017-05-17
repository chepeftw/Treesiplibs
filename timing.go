package treesiplibs

import (
    "time"
    "math/rand"
)

// Timeout functions to start and stop the timer
func StartTimeout(d int, r1 *rand.Rand) *time.Timer {
	duration := float32(r1.Intn(d*1000)/1000)
    timeout := time.NewTimer(time.Millisecond * time.Duration(1000+duration))
    return timeout
}

func StartTimeoutF(d float32) *time.Timer {
    return StartTimeoutUnited(d, 0)
}

func StartTimeoutStrong(d float32) *time.Timer {
    return StartTimeoutUnited(d, 0)
}

func StartTimeoutUnited(duration float32, offset float32) *time.Timer {
    return time.NewTimer(time.Millisecond * time.Duration( offset + duration ))
}

func StopTimeout(timer *time.Timer) {
    if timer != nil {
        timer.Stop()
    }
}