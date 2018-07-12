package golf

import (
	"time"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type JSONSuite struct{}

func (s *JSONSuite) TestJsonFloatNew(t sweet.T) {
	f := newJsonFloat(12345)

	Expect(f.val).To(Equal(float64(12345)))
}

func (s *JSONSuite) TestJsonFloatJson(t sweet.T) {
	f := newJsonFloat(float64(1440387554.671944965))

	Expect(f.val).To(Equal(1440387554.671944965))

	json, err := f.MarshalJSON()
	t.Logf("%v", string(json))
	Expect(err).To(BeNil())
	Expect(string(json)).To(Equal("1440387554.671945"))
}

func (s *JSONSuite) TestJsonNoLogger(t sweet.T) {
	msg := newMessage()
	msg.Level = LEVEL_CRIT
	msg.Hostname = "hostname"

	ts := time.Unix(0, 1440387554671944965)
	msg.Timestamp = &ts

	msg.ShortMessage = "short_message"
	msg.FullMessage = "full_message"

	msg.Attrs["attr1"] = "val1"
	msg.Attrs["attr2"] = 1234

	json, _ := generateMsgJson(msg)
	Expect(json).To(Equal(`{` +
		`"_attr1":"val1","_attr2":1234,"full_message":"full_message",` +
		`"host":"hostname","level":2,"short_message":"short_message",` +
		`"timestamp":1440387554.671945,"version":"1.1"` +
		`}`))
}

func (s *JSONSuite) TestJsonWithLogger(t sweet.T) {
	l := newLogger()
	l.SetAttr("attr1", "notval1")
	l.SetAttr("attr3", "val3")

	msg := newMessage()
	msg.logger = l
	msg.Level = LEVEL_CRIT
	msg.Hostname = "hostname"

	ts := time.Unix(0, 1440387554671944965)
	msg.Timestamp = &ts

	msg.ShortMessage = "short_message"
	msg.FullMessage = "full_message"

	msg.Attrs["attr1"] = "val1"
	msg.Attrs["attr2"] = 1234

	json, _ := generateMsgJson(msg)
	Expect(json).To(Equal(`{` +
		`"_attr1":"val1","_attr2":1234,"_attr3":"val3",` +
		`"full_message":"full_message","host":"hostname","level":2,` +
		`"short_message":"short_message","timestamp":1440387554.671945,` +
		`"version":"1.1"` +
		`}`))
}
