package clock

import (
	"fmt"
)

type Clock struct {
    hours, mins, totalMins int
}

func New(hours int, mins int) Clock {
    c := new(Clock)

    // Store total minutes so we can format for output later on.
    // Take the hours, multiple by 60 and add to the minutes as well.
    c.totalMins = (hours * 60) + mins
    return c.FormatTime()
}

// this function takes total minutes and calculates the hours & minutes
func (c Clock) FormatTime() Clock {

	c.hours = c.totalMins / 60
	c.mins = c.totalMins - (c.hours * 60)
	c.hours = c.hours % 24

	return c
}

// return the hours & minutes as a readable format HH:MM
// it will pad space with 0, so 1 hour 5 minutes will be 01:05 instead of 1:5
func (c Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", c.hours, c.mins)
}


func (c Clock) Add(addedMins int) Clock {
	c.totalMins = c.totalMins + addedMins

	return c.FormatTime()
}
