package tinyterm

import "errors"

type buffer struct {
	bytes [64]byte
	rptr  uint8
	wptr  uint8
}

func (pb *buffer) Reset() {
	pb.rptr = 0
	pb.wptr = 0
}

func (pb *buffer) WriteByte(b byte) error {
	if pb.wptr+1 < 64 {
		pb.bytes[pb.wptr] = b
		pb.wptr++
		return nil
	}
	return errBufFull
}

func (pb *buffer) ReadByte() (byte, error) {
	if pb.rptr < pb.wptr {
		pb.rptr++
		return pb.bytes[pb.rptr], nil
	}
	return 0, errBufEmpty
}

func (pb *buffer) Buffered() int {
	return int(pb.wptr - pb.rptr)
}

func (pb *buffer) IndexByte(b byte, start int) int {
	for i, v := range pb.bytes[start:] {
		if v == b {
			return i
		}
	}
	return -1
}

func (pb *buffer) Len() int {
	return int(pb.wptr)
}

func indexByte(buf []byte, b byte) int {
	for i, v := range buf {
		if v == b {
			return i
		}
	}
	return -1
}

func (pb *buffer) parseSGR() (values [5]uint8, count int) {
	buf := pb.bytes[:pb.wptr]
	for i := range values {
		if d := indexByte(buf, ';'); d > 0 {
			if v, err := parseUint8(buf[:d]); err == nil {
				values[i] = v
				buf = buf[d+1:]
				count++
				continue
			}
		} else {
			if v, err := parseUint8(buf[:]); err == nil {
				values[i] = v
				count++
			}
		}
		break
	}
	return
}

var (
	errInvalidLength    = errors.New("invalid length")
	errInvalidCharacter = errors.New("invalid character")
)

func parseUint8(buf []byte) (val uint8, err error) {
	const powers uint32 = (100 << 16) | (10 << 8) | (1)
	length := len(buf)
	if length > 0 && length < 4 {
		for i := 0; i < length; i++ {
			if c := buf[i]; c >= '0' && c <= '9' {
				val += uint8(powers>>(8*(length-i-1))) * (c - '0')
				continue
			}
			return 0, errInvalidCharacter
		}
		return val, nil
	}
	return 0, errInvalidLength
}
