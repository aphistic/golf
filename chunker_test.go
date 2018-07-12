package golf

import (
	"fmt"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type ChunkerSuite struct{}

func (s *ChunkerSuite) TestChunkerNew(t sweet.T) {
	w := newTestWriter()
	chnk, err := newChunker(w, 1234)

	Expect(err).To(BeNil())

	Expect(chnk).ToNot(BeNil())
	Expect(chnk.w).To(Equal(w))
	Expect(chnk.chunkSize).To(Equal(1234))
}

func (s *ChunkerSuite) TestChunkerNewChunkTooSmall(t sweet.T) {
	w := newTestWriter()

	_, err := newChunker(w, -1)
	Expect(err).To(Equal(ErrChunkTooSmall))
}

func (s *ChunkerSuite) TestChunkerWrite(t sweet.T) {
	chnk, _ := newChunker(nil, 15)

	Expect(chnk.buff).To(HaveLen(0))

	chnk.Write([]byte{1, 2, 3, 4})
	Expect(chnk.buff).To(Equal([]byte{1, 2, 3, 4}))

	chnk.Write([]byte{5, 6, 7, 8})
	Expect(chnk.buff).To(Equal([]byte{1, 2, 3, 4, 5, 6, 7, 8}))
}

func (s *ChunkerSuite) TestChunkerFlush(t sweet.T) {
	// With a chunk size of 13 it should leave space for 1 byte of
	// data after the header (12 bytes). So there should be a total
	// of 5 packets written

	w := newTestWriter()
	chnk, _ := newChunker(w, 13)

	chnk.Write([]byte{1, 2, 3, 4, 5})
	Expect(chnk.buff).To(Equal([]byte{1, 2, 3, 4, 5}))

	chnk.flushWithId([]byte{1, 2, 3, 4, 5, 6, 7, 8})

	Expect(w.Written).To(HaveLen(5))

	Expect(w.Written[0]).To(Equal([]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 0, 5, 1}))
	Expect(w.Written[1]).To(Equal([]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 1, 5, 2}))
	Expect(w.Written[2]).To(Equal([]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 2, 5, 3}))
	Expect(w.Written[3]).To(Equal([]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 3, 5, 4}))
	Expect(w.Written[4]).To(Equal([]byte{0x1e, 0x0f, 1, 2, 3, 4, 5, 6, 7, 8, 4, 5, 5}))

	Expect(chnk.buff).To(Equal([]byte{}))
}

func (s *ChunkerSuite) TestChunkerFlushIdLength(t sweet.T) {
	chnk, _ := newChunker(newTestWriter(), 13)
	chnk.Write([]byte{1, 2, 3, 4, 5})

	err := chnk.flushWithId([]byte{1, 2, 3, 4, 5, 6, 7})
	Expect(err).To(Equal(fmt.Errorf("id length must be equal to 8")))

	err = chnk.flushWithId([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	Expect(err).To(Equal(fmt.Errorf("id length must be equal to 8")))
}
