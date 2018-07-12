package golf

import (
	"errors"
)

var (
	ErrChunkTooSmall = errors.New("chunk size is too small, it must be at least 13")
)
