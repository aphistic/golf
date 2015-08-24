package golf

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type GolfSuite struct{}

var _ = Suite(&GolfSuite{})

func Example() {
	c, _ := NewClient()
	c.Dial("udp://localhost")

	l, _ := c.NewLogger()
	l.SetAttr("facility", "example.facility")

	wait := make(chan int)
	go func() {
		for idx := 1; idx <= 10; idx++ {
			l.Dbgm(
				map[string]interface{}{
					"attr1": "val1",
					"attr2": 1234},
				"Test %v",
				idx)
			idx += 1
		}
		wait <- 1
	}()
	<-wait
}

func Example_defaultLogger() {
	c, _ := NewClient()
	c.Dial("udp://localhost")

	l, _ := c.NewLogger()
	l.SetAttr("facility", "example.facility")
	DefaultLogger(l)

	wait := make(chan int)
	go func() {
		for idx := 1; idx <= 10; idx++ {
			Dbgm(
				map[string]interface{}{
					"attr1": "val1",
					"attr2": 1234},
				"Test %v",
				idx)
			idx += 1
		}
		wait <- 1
	}()
	<-wait
}
