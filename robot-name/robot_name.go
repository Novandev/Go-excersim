package robotname

import (
	"math/rand"
	"fmt"
	"time"
)

//Declare a robot as a string name
type Robot string

//An array that has all of the names currentl in use by making a map of bools
var namesInUse = make(map[Robot]bool)

// interface that returns a robots name
func (r *Robot) Name() string {
	rand.Seed(time.Now().UnixNano())

	for inUse, ok := string(*r) == "", true; ok && inUse; inUse, ok = namesInUse[*r] {
		*r = Robot(letter() + letter() + number())
	}
	namesInUse[*r] = true
	return string(*r)
}

// interfaces that modifyes the robot and gives it a new string name
func (r *Robot) Reset() {
	*r = Robot("")
}
// Number generator fot adding digits to a tring robot name ex ZUI235
func number() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

// Genetatior funtion that gets a radom uppercase letter
func letter() string {
	return string(rand.Intn('Z'-'A') + 'A')
}

