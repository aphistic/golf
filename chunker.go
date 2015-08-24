package golf

import (
	"github.com/satori/go.uuid"
	"io"
	"math"
)

type chunker struct {
	chunkSize int
	buff      []byte
	w         io.Writer
}

func newChunker(w io.Writer, chunkSize int) *chunker {
	return &chunker{
		chunkSize: chunkSize,
		buff:      make([]byte, 0),
		w:         w,
	}
}

func (c *chunker) Write(p []byte) (int, error) {
	c.buff = append(c.buff, p...)
	return len(p), nil
}

func (c *chunker) Flush() error {
	offset := 0
	buffLen := len(c.buff)
	chunkSize := c.chunkSize - 12

	idFull := uuid.NewV4()

	chunkBuff := make([]byte, c.chunkSize)
	copy(chunkBuff[0:2], []byte{0x1e, 0x0f})
	copy(chunkBuff[2:10], idFull.Bytes())

	totalChunks := int(math.Ceil(float64(buffLen) / float64(chunkSize)))
	chunkBuff[11] = byte(totalChunks)

	for {
		//fmt.Printf("%v/%v\n", chunkIdx, totalChunks)
		left := buffLen - offset
		if left > chunkSize {
			copy(chunkBuff[12:], c.buff[offset:offset+chunkSize])
			c.w.Write(chunkBuff)
		} else {
			copy(chunkBuff[12:], c.buff[offset:offset+left])
			c.w.Write(chunkBuff[0 : left+12])
			break
		}

		offset += chunkSize
		chunkBuff[10] += 1
	}

	return nil
}
