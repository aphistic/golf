package golf

import (
	. "gopkg.in/check.v1"
)

func (s *GolfSuite) TestGenDefaultMessageWithFormatNoParams(c *C) {
	cl, err := NewClient()
	c.Assert(cl, NotNil)
	c.Assert(err, IsNil)
	l, err := cl.NewLogger()
	c.Assert(l, NotNil)
	c.Assert(err, IsNil)
	DefaultLogger(l)

	// Tests to make sure a string won't be double-formatted if
	// no paramters are passed to the format string
	msg := genDefaultMsg(nil, 1, "%2b")
	c.Check(msg.Level, Equals, 1)
	c.Check(msg.ShortMessage, Equals, "%2b")
	c.Check(msg.Attrs, IsNil)
}

func (s *GolfSuite) TestGenDefaultMessageWithFormat(c *C) {
	cl, err := NewClient()
	c.Assert(cl, NotNil)
	c.Assert(err, IsNil)
	l, err := cl.NewLogger()
	c.Assert(l, NotNil)
	c.Assert(err, IsNil)
	DefaultLogger(l)

	// Tests to make sure a string will be formatted if
	// paramters are passed to the format string
	msg := genDefaultMsg(nil, 1, "%2b", true)
	c.Check(msg.Level, Equals, 1)
	c.Check(msg.ShortMessage, Equals, "%!b(bool=true)")
	c.Check(msg.Attrs, IsNil)
}
