package clock

import "fmt"
type Clock struct{
	hours,minutes int
}


func New(hours, minutes int)*Clock{
	clock := new(Clock)
	clock.hours = hours % 24
	clock.minutes = minutes % 60
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
	clock.hours = (clock.hours + (minutes/60)) % 24
	clock.minutes = (clock.minutes + (minutes%60))
	fmt.Println(clock.hours)
}
func (clock *Clock) subtract(minutes int){
	clock.hours = (clock.hours - (minutes/60)) % 24
	clock.minutes = (clock.minutes - (minutes%60))
	if clock.minutes < 0{
		clock.minutes = -clock.minutes
	}
	if clock.hours < 0{
		clock.hours = -clock.hours
	}
}
