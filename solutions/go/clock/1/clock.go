package clock

import "fmt"

type Clock struct {
    hour   int
    minute int
}

func normalize(h, m int) (int, int) {
	totalMinutes := h*60 + m
	totalMinutes = ((totalMinutes % 1440) + 1440) % 1440
	return totalMinutes / 60, totalMinutes % 60
}

func New(h, m int) Clock {
	h, m = normalize(h, m)
	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}