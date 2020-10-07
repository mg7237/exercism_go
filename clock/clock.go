package clock

import (
	"fmt"
)

//Clock struct defines a clock consisting of hour and minute
type Clock struct {
	minute int
}

// New - constructor for clock
func New(h int, m int) Clock {
	return buildClock(h*60 + m)
}

// Add adds minute to a clock struct
func (c Clock) Add(m int) Clock {
	return buildClock(c.minute + m)
}

// Subtract subtracts minute from clock struct
func (c Clock) Subtract(m int) Clock {
	return buildClock(c.minute - m)
}

func buildClock(minutes int) Clock {
	hours := int(minutes / 60)
	if minutes < 0 && minutes%60 != 0 {
		hours--
	}

	minutes -= hours * 60
	carryDays := int(hours / 24)
	if hours < 0 && hours%24 != 0 {
		carryDays--
	}
	hours -= (carryDays * 24)
	return Clock{hours*60 + minutes}
}

// Clock stringer
func (c Clock) String() string {
	hours := int(c.minute / 60)
	minutes := c.minute - (hours * 60)
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
