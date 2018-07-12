package golf

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

func (s *GolfSuite) TestGenDefaultMessageWithFormatNoParams(t sweet.T) {
	cl, err := NewClient()
	Expect(err).To(BeNil())
	Expect(cl).ToNot(BeNil())

	l, err := cl.NewLogger()
	Expect(err).To(BeNil())
	Expect(l).ToNot(BeNil())

	DefaultLogger(l)

	// Tests to make sure a string won't be double-formatted if
	// no paramters are passed to the format string
	msg := genDefaultMsg(nil, 1, "%2b")
	Expect(msg.Level).To(Equal(1))
	Expect(msg.ShortMessage).To(Equal("%2b"))
	Expect(msg.Attrs).To(BeNil())
}

func (s *GolfSuite) TestGenDefaultMessageWithFormat(t sweet.T) {
	cl, err := NewClient()
	Expect(err).To(BeNil())
	Expect(cl).ToNot(BeNil())

	l, err := cl.NewLogger()
	Expect(err).To(BeNil())
	Expect(l).ToNot(BeNil())

	DefaultLogger(l)

	// Tests to make sure a string will be formatted if
	// paramters are passed to the format string
	msg := genDefaultMsg(nil, 1, "%2b", true)
	Expect(msg.Level).To(Equal(1))
	Expect(msg.ShortMessage).To(Equal("%!b(bool=true)"))
	Expect(msg.Attrs).To(BeNil())
}
