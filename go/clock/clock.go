package clock

import "fmt"

type Clock int

// Create and return a clock based on the hours and minutes given
func New(hours, mins int) Clock {

	clock := Clock( (hours * 60 + mins) % (24 * 60) )

	// if the clock is less than 0, add a days worth of minutes to it
	// this is because both the hours and minutes passed in can be more than and less than a day (25:00 or -25:00 etc)
	for clock < 0 {
		clock += (24 * 60)
	}

	return clock
}

// return the clock as a readable format HH:MM
// it will pad space with 0, so 1 hour 5 minutes will be 01:05 instead of 1:5
func (clock Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", clock/60, clock%60)
}

// Add minutes to the clock we already have - because it's done via New() it'll handle sorting the hours etc
func (clock Clock) Add(mins int) Clock {
	return New(0, int(clock) + mins)
}