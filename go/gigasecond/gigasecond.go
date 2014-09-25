package gigasecond

import "time"

// Gigasecond is 1 billion seconds (10^9)
const Gigasecond = 1e9 * time.Second

// AddGigasecond adds a gigasecond to the time given
func AddGigasecond(input time.Time) time.Time {
	return input.Add(Gigasecond)
}
