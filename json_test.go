package golf

import (
	. "gopkg.in/check.v1"
	"time"
)

func (s *GolfSuite) TestJsonFloatNew(c *C) {
	f := newJsonFloat(12345)
	c.Check(f.val, Equals, float64(12345))
}

func (s *GolfSuite) TestJsonFloatJson(c *C) {
	f := newJsonFloat(float64(1440387554.671944965))

	c.Check(f.val, Equals, 1440387554.671944965)

	json, err := f.MarshalJSON()
	c.Logf("%v", string(json))
	c.Check(err, IsNil)
	c.Check(string(json), Equals, "1440387554.671945")
}

func (s *GolfSuite) TestJsonNoLogger(c *C) {
	msg := newMessage()
	msg.Level = LEVEL_CRIT
	msg.Hostname = "hostname"

	t := time.Unix(0, 1440387554671944965)
	msg.Timestamp = &t

	msg.ShortMessage = "short_message"
	msg.FullMessage = "full_message"

	msg.Attrs["attr1"] = "val1"
	msg.Attrs["attr2"] = 1234

	json, _ := generateMsgJson(msg)
	c.Check(json, Equals, `{`+
		`"_attr1":"val1","_attr2":1234,"full_message":"full_message",`+
		`"host":"hostname","level":2,"short_message":"short_message",`+
		`"timestamp":1440387554.671945,"version":"1.1"`+
		`}`)
}

func (s *GolfSuite) TestJsonWithLogger(c *C) {
	l := newLogger()
	l.SetAttr("attr1", "notval1")
	l.SetAttr("attr3", "val3")

	msg := newMessage()
	msg.logger = l
	msg.Level = LEVEL_CRIT
	msg.Hostname = "hostname"

	t := time.Unix(0, 1440387554671944965)
	msg.Timestamp = &t

	msg.ShortMessage = "short_message"
	msg.FullMessage = "full_message"

	msg.Attrs["attr1"] = "val1"
	msg.Attrs["attr2"] = 1234

	json, _ := generateMsgJson(msg)
	c.Check(json, Equals, `{`+
		`"_attr1":"val1","_attr2":1234,"_attr3":"val3",`+
		`"full_message":"full_message","host":"hostname","level":2,`+
		`"short_message":"short_message","timestamp":1440387554.671945,`+
		`"version":"1.1"`+
		`}`)
}
