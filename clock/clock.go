package clock

import "fmt"

type Clock struct {
	hour, minute uint
}

func New(hour int, minute int) Clock  {
	appendHour := minute / 60

	if minute = minute % 60; minute < 0 {
		minute += 60
		appendHour -= 1
	}

	hour = hour + appendHour

	if hour = hour % 24; hour < 0 {
		hour += 24
	}

	return Clock{uint(hour), uint(minute)}
	//return c
}

func (c Clock) String() string  {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minute int) Clock {
	m := int(c.minute) + minute

	appendHour := m / 60

	m = m % 60

	h := int(c.hour) + appendHour

	h = h % 24


	c.hour, c.minute = uint(h), uint(m)
	return c
}


func (c Clock) Subtract(minute int) Clock {
	m := int(c.minute) - minute

	appendHour := m / 60

	m = m % 60

	if m = m % 60; m < 0 {
		m = m + 60
		appendHour -= 1
	}

	h := int(c.hour) + appendHour

	if h = h % 24; h < 0 {
		h = h + 24
	}


	c.hour, c.minute = uint(h), uint(m)
	return c
}
