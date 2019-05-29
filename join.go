package berry

import (
	"bytes"
)

var sseq = []uint8{'\x1b', '['}

func join(codes []uint8) []uint8 {
	if bytes.HasPrefix(codes, sseq) {
		return codes
	}

	buf := make([]byte, len(codes)*4+1)
	n := 0

	copy(buf[n:], sseq)
	n += 2

	for _, c := range codes {
		if c < 10 {
			buf[n] = one[c]
			n++
		} else if c < 100 {
			copy(buf[n:], two[(c-10)*2:(c-9)*2])
			n += 2
		} else {
			cc := uint16(c)
			copy(
				buf[n:],
				three[(cc-100)*3:(cc-99)*3],
			)
			n += 3
		}
		buf[n] = ';'
		n++
	}

	buf[n-1] = 'm'

	return buf[:n]
}

var (
	// ruby: (0..9).map(&:join).map {|str| str.split("").map(&:ord)}
	one = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

	// ruby: (10..99).each_slice(10).map(&:join).map {|str| str.split("").map(&:ord)}
	two = []byte{49, 48, 49, 49, 49, 50, 49, 51, 49, 52, 49, 53, 49, 54, 49, 55, 49, 56, 49, 57,
		50, 48, 50, 49, 50, 50, 50, 51, 50, 52, 50, 53, 50, 54, 50, 55, 50, 56, 50, 57,
		51, 48, 51, 49, 51, 50, 51, 51, 51, 52, 51, 53, 51, 54, 51, 55, 51, 56, 51, 57,
		52, 48, 52, 49, 52, 50, 52, 51, 52, 52, 52, 53, 52, 54, 52, 55, 52, 56, 52, 57,
		53, 48, 53, 49, 53, 50, 53, 51, 53, 52, 53, 53, 53, 54, 53, 55, 53, 56, 53, 57,
		54, 48, 54, 49, 54, 50, 54, 51, 54, 52, 54, 53, 54, 54, 54, 55, 54, 56, 54, 57,
		55, 48, 55, 49, 55, 50, 55, 51, 55, 52, 55, 53, 55, 54, 55, 55, 55, 56, 55, 57,
		56, 48, 56, 49, 56, 50, 56, 51, 56, 52, 56, 53, 56, 54, 56, 55, 56, 56, 56, 57,
		57, 48, 57, 49, 57, 50, 57, 51, 57, 52, 57, 53, 57, 54, 57, 55, 57, 56, 57, 57,
	}

	// ruby: (100..255).each_slice(10).map(&:join).map {|str| str.split("").map(&:ord)}
	three = []byte{49, 48, 48, 49, 48, 49, 49, 48, 50, 49, 48, 51, 49, 48, 52, 49, 48, 53, 49, 48, 54, 49, 48, 55, 49, 48, 56, 49, 48, 57,
		49, 49, 48, 49, 49, 49, 49, 49, 50, 49, 49, 51, 49, 49, 52, 49, 49, 53, 49, 49, 54, 49, 49, 55, 49, 49, 56, 49, 49, 57,
		49, 50, 48, 49, 50, 49, 49, 50, 50, 49, 50, 51, 49, 50, 52, 49, 50, 53, 49, 50, 54, 49, 50, 55, 49, 50, 56, 49, 50, 57,
		49, 51, 48, 49, 51, 49, 49, 51, 50, 49, 51, 51, 49, 51, 52, 49, 51, 53, 49, 51, 54, 49, 51, 55, 49, 51, 56, 49, 51, 57,
		49, 52, 48, 49, 52, 49, 49, 52, 50, 49, 52, 51, 49, 52, 52, 49, 52, 53, 49, 52, 54, 49, 52, 55, 49, 52, 56, 49, 52, 57,
		49, 53, 48, 49, 53, 49, 49, 53, 50, 49, 53, 51, 49, 53, 52, 49, 53, 53, 49, 53, 54, 49, 53, 55, 49, 53, 56, 49, 53, 57,
		49, 54, 48, 49, 54, 49, 49, 54, 50, 49, 54, 51, 49, 54, 52, 49, 54, 53, 49, 54, 54, 49, 54, 55, 49, 54, 56, 49, 54, 57,
		49, 55, 48, 49, 55, 49, 49, 55, 50, 49, 55, 51, 49, 55, 52, 49, 55, 53, 49, 55, 54, 49, 55, 55, 49, 55, 56, 49, 55, 57,
		49, 56, 48, 49, 56, 49, 49, 56, 50, 49, 56, 51, 49, 56, 52, 49, 56, 53, 49, 56, 54, 49, 56, 55, 49, 56, 56, 49, 56, 57,
		49, 57, 48, 49, 57, 49, 49, 57, 50, 49, 57, 51, 49, 57, 52, 49, 57, 53, 49, 57, 54, 49, 57, 55, 49, 57, 56, 49, 57, 57,
		50, 48, 48, 50, 48, 49, 50, 48, 50, 50, 48, 51, 50, 48, 52, 50, 48, 53, 50, 48, 54, 50, 48, 55, 50, 48, 56, 50, 48, 57,
		50, 49, 48, 50, 49, 49, 50, 49, 50, 50, 49, 51, 50, 49, 52, 50, 49, 53, 50, 49, 54, 50, 49, 55, 50, 49, 56, 50, 49, 57,
		50, 50, 48, 50, 50, 49, 50, 50, 50, 50, 50, 51, 50, 50, 52, 50, 50, 53, 50, 50, 54, 50, 50, 55, 50, 50, 56, 50, 50, 57,
		50, 51, 48, 50, 51, 49, 50, 51, 50, 50, 51, 51, 50, 51, 52, 50, 51, 53, 50, 51, 54, 50, 51, 55, 50, 51, 56, 50, 51, 57,
		50, 52, 48, 50, 52, 49, 50, 52, 50, 50, 52, 51, 50, 52, 52, 50, 52, 53, 50, 52, 54, 50, 52, 55, 50, 52, 56, 50, 52, 57,
		50, 53, 48, 50, 53, 49, 50, 53, 50, 50, 53, 51, 50, 53, 52, 50, 53, 53}
)
