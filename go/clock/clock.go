package clock

import (
	"fmt"
)

// Clock implements a clock structure
type Clock struct {
	Hours   int
	Minutes int
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)

}

func roll(hours, min int) (sumhour int, summin int) {

	min += hours * 60
	sumhour = min / 60
	summin = min % 60
	if summin < 0 {
		summin += 60
		sumhour--
	}
	sumhour = sumhour % 24
	if sumhour < 0 {
		sumhour += 24
	}

	return sumhour, summin

}

// Add adds given minutes to clock
func (c Clock) Add(minutes int) Clock {
	c.Hours, c.Minutes = roll(c.Hours, c.Minutes+minutes)
	return c
}

// Subtract removes given minutes to clock
func (c Clock) Subtract(minutes int) Clock {
	c.Hours, c.Minutes = roll(c.Hours, c.Minutes-minutes)
	return c
}

// New creates a new clock
func New(hour int, minute int) Clock {

	hour, minute = roll(hour, minute)
	return Clock{hour, minute}

}
