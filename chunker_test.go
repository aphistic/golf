package golf

import (
	. "gopkg.in/check.v1"
)

func (s *GolfSuite) TestChunkerNew(c *C) {
	w := newTestWriter()
	chnk, err := newChunker(w, 1234)

	c.Check(err, IsNil)

	c.Check(chnk, NotNil)
	c.Check(chnk, NotNil)
	c.Check(chnk.w, Equals, w)
	c.Check(chnk.chunkSize, Equals, 1234)
}

func (s *GolfSuite) TestChunkerNewChunkTooSmall(c *C) {
	w := newTestWriter()

	_, err := newChunker(w, -1)
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "Chunk size must be at least 13.")
}

func (s *GolfSuite) TestChunkerWrite(c *C) {
	chnk, _ := newChunker(nil, 15)

	c.Check(chnk.buff, HasLen, 0)

	chnk.Write([]byte{1, 2, 3, 4})
	c.Check(chnk.buff, HasLen, 4)
	c.Check(chnk.buff[0:], DeepEquals, []byte{1, 2, 3, 4})

	chnk.Write([]byte{5, 6, 7, 8})
	c.Check(chnk.buff, HasLen, 8)
	c.Check(chnk.buff[0:], DeepEquals, []byte{1, 2, 3, 4, 5, 6, 7, 8})
}

func (s *GolfSuite) TestChunkerFlush(c *C) {
	// With a chunk size of 13 it should leave space for 1 byte of
	// data after the header (12 bytes). So there should be a total
	// of 5 packets written

	w := newTestWriter()
	chnk, _ := newChunker(w, 13)

	chnk.Write([]byte{1, 2, 3, 4, 5})
	c.Check(chnk.buff, HasLen, 5)

	chnk.flushWithId([]byte{1, 2, 3, 4, 5, 6, 7, 8})

	c.Assert(w.Written, HasLen, 5)

	c.Check(w.Written[0], DeepEquals,
		[]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 0, 5, 1})
	c.Check(w.Written[1], DeepEquals,
		[]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 1, 5, 2})
	c.Check(w.Written[2], DeepEquals,
		[]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 2, 5, 3})
	c.Check(w.Written[3], DeepEquals,
		[]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 3, 5, 4})
	c.Check(w.Written[4], DeepEquals,
		[]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 4, 5, 5})

	c.Check(chnk.buff, HasLen, 0)
}

func (s *GolfSuite) TestChunkerFlushIdLength(c *C) {
	chnk, _ := newChunker(newTestWriter(), 13)
	chnk.Write([]byte{1, 2, 3, 4, 5})

	err := chnk.flushWithId([]byte{1, 2, 3, 4, 5, 6, 7})
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "id length must be equal to 8")

	err = chnk.flushWithId([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "id length must be equal to 8")
}
