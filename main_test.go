package golf

import (
	"testing"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	RegisterFailHandler(sweet.GomegaFail)

	sweet.Run(m, func(s *sweet.S) {
		s.AddSuite(&ChunkerSuite{})
		s.AddSuite(&GolfSuite{})
	})
}

type GolfSuite struct{}

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

type testWriter struct {
	Written [][]byte
}

func newTestWriter() *testWriter {
	return &testWriter{
		Written: make([][]byte, 0),
	}
}
func (tw *testWriter) reset() {
	tw.Written = make([][]byte, 0)
}
func (tw *testWriter) Write(p []byte) (int, error) {
	// Need to copy it because of the way the chunker reuses the
	// write buffer.  That may be a problem at some point in the future...
	// Need to keep an eye on that.
	write := make([]byte, len(p))
	copy(write, p)
	tw.Written = append(tw.Written, write)

	return len(p), nil
}
