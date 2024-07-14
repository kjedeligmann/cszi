package main

func (s *bit8) IP() {
	// permutation is {1, 5, 2, 0, 3, 7, 4, 6}
	*s = (*s << 1 & (1 << 7) >> 0) |
		(*s << 5 & (1 << 7) >> 1) |
		(*s << 2 & (1 << 7) >> 2) |
		(*s << 0 & (1 << 7) >> 3) |
		(*s << 3 & (1 << 7) >> 4) |
		(*s << 7 & (1 << 7) >> 5) |
		(*s << 4 & (1 << 7) >> 6) |
		(*s << 6 & (1 << 7) >> 7)
}

func (s *bit8) IPr() {
	// permutation is {3, 0, 2, 4, 6, 1, 7, 5}
	*s = (*s << 3 & (1 << 7) >> 0) |
		(*s << 0 & (1 << 7) >> 1) |
		(*s << 2 & (1 << 7) >> 2) |
		(*s << 4 & (1 << 7) >> 3) |
		(*s << 6 & (1 << 7) >> 4) |
		(*s << 1 & (1 << 7) >> 5) |
		(*s << 7 & (1 << 7) >> 6) |
		(*s << 5 & (1 << 7) >> 7)
}

func (s *bit4) P4() {
	// permutation is {1, 3, 2, 0}
	*s = (*s << 1 & (1 << 3) >> 0) |
		(*s << 3 & (1 << 3) >> 1) |
		(*s << 2 & (1 << 3) >> 2) |
		(*s << 0 & (1 << 3) >> 3)
}

func (s *bit8) SW() {
	*s = *s<<4 | *s>>4
}

func Sb(s bit8) (result bit4) {
	// A function that implements the use of substitution matrices
	var S0 uint32 = 0b01001110111001000010011111011110
	var S1 uint32 = 0b00011011100001111100010010010011
	var row0, row1, col0, col1 uint8
	col0 = uint8(s>>5) & 0b11
	col1 = uint8(s>>1) & 0b11
	row0 = (uint8(s>>7<<1) | uint8(s<<3>>7)) & 0b11
	row1 = (uint8(s>>3<<1) | uint8(s<<7>>7)) & 0b11
	result = bit4(
		(S0 >> (30 - (row0 << 3) - (col0 << 1)) << 2) |
			(S1 >> (30 - (row1 << 3) - (col1 << 1)) & 0b11))
	result.Mask()
	return
}

func EP(s bit4) (result bit8) {
	result |= bit8(s) << 7
	result |= bit8(s) >> 1 << 4
	result |= bit8(s) << 5 >> 4
	result |= bit8(s) >> 3
	return
}

func SDESe(plaintext string, key bit10) string {
	k1, k2 := KeyGen(key)
    ciphertext := make([]byte, len(plaintext))

	for i, b := range []byte(plaintext) {
		block := bit8(b)
		block.IP()

		var left bit4 = bit4(block >> 4)
		var right bit4 = bit4(block)
		right.Mask()

		expanded := EP(right)
		expanded ^= k1
		sboxed := Sb(expanded)
		sboxed.P4()

		left ^= sboxed

		block = bit8(left<<4) | bit8(right)
		block.SW()

		var left2 bit4 = bit4(block >> 4)
		var right2 bit4 = bit4(block)
		right2.Mask()

		expanded2 := EP(right2)
		expanded2 ^= k2
		sboxed2 := Sb(expanded2)
		sboxed2.P4()

		left2 ^= sboxed2

		block = bit8(left2<<4) | bit8(right2)
		block.IPr()

		ciphertext[i] = byte(block)
	}

	return string(ciphertext)
}

func SDESd(ciphertext string, key bit10) string {
	k1, k2 := KeyGen(key)
    plaintext := make([]byte, len(ciphertext))

	for i, b := range []byte(ciphertext) {
		block := bit8(b)
		block.IP()

		var left bit4 = bit4(block >> 4)
		var right bit4 = bit4(block)
		right.Mask()

		expanded := EP(right)
		expanded ^= k2
		sboxed := Sb(expanded)
		sboxed.P4()

		left ^= sboxed

		block = bit8(left<<4) | bit8(right)
		block.SW()

		var left2 bit4 = bit4(block >> 4)
		var right2 bit4 = bit4(block)
		right2.Mask()

		expanded2 := EP(right2)
		expanded2 ^= k1
		sboxed2 := Sb(expanded2)
		sboxed2.P4()

		left2 ^= sboxed2

		block = bit8(left2<<4) | bit8(right2)
		block.IPr()

		plaintext[i] = byte(block)
	}

	return string(plaintext)
}
