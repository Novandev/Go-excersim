package clock

import "fmt"

type Clock struct{
	hours,minutes int
}


func New(hours, minutes int)*Clock{
	clock := new(Clock)
	clock.hours = hours
	clock.minutes = minutes
	return clock
}
// {8, 0, "08:00"}
func (clock *Clock) String() string {
	returned_clock :=""
	if clock.hours < 10 && clock.minutes > 10{
		returned_clock = fmt.Sprintf("0%d:%d", clock.hours, clock.minutes)
	}
	if clock.hours < 10 && clock.minutes < 10{
		returned_clock = fmt.Sprintf("0%d:0%d", clock.hours, clock.minutes)
	}
	if clock.hours > 10 && clock.minutes < 10{
		returned_clock =fmt.Sprintf("%d:0%d", clock.hours, clock.minutes)
	}
	if clock.hours > 10 && clock.minutes > 10{
		returned_clock =fmt.Sprintf("%d:%d", clock.hours, clock.minutes)
	}

	return returned_clock
}


func (clock *Clock) add(minutes int){
	for i := 0; i < minutes; i++{
		clock.minutes += i
		if clock.minutes > 60{
			clock.hours++
			if clock.hours > 24{
				clock.hours = 0
			}
			clock.minutes = 0
		}
	}
}

// func (clock *Clock) subtract(){

// }