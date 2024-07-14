package main

type bit4 uint8
type bit5 uint8
type bit8 uint8
type bit10 uint16

func (s *bit4) Mask() {
	*s &= 0b1111
}

func (s *bit5) Mask() {
	*s &= 0b11111
}

func (s *bit10) Mask() {
	*s &= 0b1111111111
}

func (s *bit5) LS1() {
	*s = *s<<1 | *s>>4
	s.Mask()
}

func (s *bit5) LS2() {
	*s = *s<<2 | *s>>3
	s.Mask()
}

func P10(s bit10) (fi, se bit5) {
	// permutation is {6, 2, 0, 4, 1, 9, 3, 8, 7, 5}

	var sp bit10
	sp = (s << 6 & (1 << 9) >> 0) |
		(s << 2 & (1 << 9) >> 1) |
		(s << 0 & (1 << 9) >> 2) |
		(s << 4 & (1 << 9) >> 3) |
		(s << 1 & (1 << 9) >> 4) |
		(s << 9 & (1 << 9) >> 5) |
		(s << 3 & (1 << 9) >> 6) |
		(s << 8 & (1 << 9) >> 7) |
		(s << 7 & (1 << 9) >> 8) |
		(s << 5 & (1 << 9) >> 9)

	fi = bit5(sp >> 5)
	se = bit5(sp)
	fi.Mask()
	se.Mask()
	return
}

func P8(fi, se bit5) bit8 {
	// permutation is {5, 2, 6, 3, 7, 4, 9, 8}

	var s bit10
	s += bit10(fi << 5)
	s += bit10(se)

	return bit8(((s << 5 & (1 << 9) >> 0) |
		(s << 2 & (1 << 9) >> 1) |
		(s << 6 & (1 << 9) >> 2) |
		(s << 3 & (1 << 9) >> 3) |
		(s << 7 & (1 << 9) >> 4) |
		(s << 4 & (1 << 9) >> 5) |
		(s << 9 & (1 << 9) >> 6) |
		(s << 8 & (1 << 9) >> 7)) >> 2)
}

func KeyGen(k bit10) (k1, k2 bit8) {
	k.Mask()
	first, second := P10(k)
	first.LS1()
	second.LS1()
	k1 = P8(first, second)
	first.LS2()
	second.LS2()
	k2 = P8(first, second)
	return
}
