package main

import "fmt"

func main() {
	key := [8]uint32{
		0x74686973,
		0x5f69735f,
		0x615f7061,
		0x73775f66,
		0x6f725f47,
		0x4f53545f,
		0x32383134,
		0x375f3839,
	}

	fmt.Println("Режим простої заміни:")
	plaintext := "hello, world!"
	fmt.Println("Повідомлення:", plaintext)
	ciphertext := GOSTpze(plaintext, key)
	fmt.Println("Шифротекст:  ", ciphertext)
	decrypted := GOSTpzd(ciphertext, key)
	fmt.Println("Розшифровано:", decrypted)

	fmt.Printf("Ключ: %08x\n", key)
	fmt.Printf("Повідомлення: %08x\n", plaintext)
	fmt.Printf("Шифротекст:   %08x\n", ciphertext)
	fmt.Printf("Розшифрований шифротекст: %08x\n", decrypted)

	fmt.Println("Режим гамування:")
	message := "hello, world!"
	var iv uint64 = 0x1110001199999999
	fmt.Println("Повідомлення:", message)
	gammed := GOSTrg(message, key, iv)
	fmt.Printf("Hex повідомлення:                    %x\n", message)
	fmt.Printf("Результат роботи в режимі гамування: %x\n", gammed)
	gammed = GOSTrg(gammed, key, iv)
	fmt.Printf("Результат дешифрування:              %x\n", gammed)

	fmt.Println("Режим гамування зі зворотнім зв'язком:")
	fmt.Println("Повідомлення:", message)
	fmt.Printf("Hex повідомлення:                                          %x\n", message)
	gammed2 := GOSTrgzze(message, key, iv)
	fmt.Printf("Результат роботи в режимі гамування зі зворотнім зв'язком: %x\n", gammed2)
	gammed2 = GOSTrgzzd(gammed2, key, iv)
	fmt.Printf("Результат дешифрування:                                    %x\n", gammed2)
}

// id-Gost28147-89-CryptoPro-D-ParamSet or E-D
var Sboxes = []uint64{
	0xFC2A645079ED1B83,
	0xB634CFE27D805A91,
	0x1CB0FE65AD489372,
	0x15ECA70D62B493F8,
	0x0C89D2AB73654EF1,
	0x80F325EB1A47C9D6,
	0x306F1E92D8C4BA57,
	0x1A68FB04C3597D2E,
}

func Sboxed(h uint32) (result uint32) {
	var mask uint32 = 0xF0000000
	for i := 0; i < 8; i++ {
		var n uint32 = (h & mask) >> uint32(4*(7-i))
		result |= uint32(Sboxes[i]<<(n*4)>>(i*4+32)) & mask
		mask >>= 4
	}
	return
}

func Round(block uint64, key uint32) (result uint64) {
	result = block << 32
	right := uint32(block)
	right += key
	right = Sboxed(right)
	right = right<<11 | right>>21
	right ^= uint32(block >> 32)
	result |= uint64(right)
	return
}

func GOSTpze(text string, key [8]uint32) string {
	input := []byte(text)
	result := []byte{}

	if r := byte(8 - (len(input) % 8)); r != 0 && r != 8 {
		for i := 0; i < int(r); i++ {
			input = append(input, r)
		}
	}

	for i := 0; i < len(input); i += 8 {
		var block uint64
		for j := 0; j < 8; j++ {
			block |= uint64(input[i+j]) << uint64(8*(7-j))
		}

		for j := 0; j < 24; j++ {
			block = Round(block, key[j%8])
		}
		for j := 7; j >= 0; j-- {
			block = Round(block, key[j])
		}
		block = block<<32 | block>>32

		for j := 0; j < 8; j++ {
			result = append(result, byte(block>>uint64(8*(7-j))))
		}
	}

	return string(result)
}

func GOSTpzd(text string, key [8]uint32) string {
	input := []byte(text)
	result := []byte{}

	for i := 0; i < len(input); i += 8 {
		var block uint64
		for j := 0; j < 8; j++ {
			block |= uint64(input[i+j]) << uint64(8*(7-j))
		}

		for j := 0; j < 8; j++ {
			block = Round(block, key[j])
		}
		for j := 23; j >= 0; j-- {
			block = Round(block, key[j%8])
		}
		block = block<<32 | block>>32

		for j := 0; j < 8; j++ {
			result = append(result, byte(block>>uint64(8*(7-j))))
		}
	}

	r := result[len(result)-1]
	if int(r) < 8 {
		var isPadded bool
		for i := 0; i < int(r); i++ {
			if result[len(result)-i-1] != r {
				isPadded = false
				break
			}
		}
		if isPadded {
			result = result[:r]
		}
	}
	return string(result)
}

func GOSTrg(text string, key [8]uint32, initVector uint64) string {
	initVectorBytes := []byte{}
	for i := 0; i < 8; i++ {
		initVectorBytes = append(initVectorBytes, byte(initVector>>uint64(8*(7-i))))
	}
	initVectorBytes = []byte(GOSTpze(string(initVectorBytes), key))

	var left, right uint32
	for i := 0; i < 4; i++ {
		left |= uint32(initVectorBytes[i]) << uint32(8*(3-i))
		right |= uint32(initVectorBytes[i+4]) << uint32(8*(3-i))
	}

	const c1, c2 uint32 = 0x1010104, 0x1010101
	input := []byte(text)
	var gammaInt uint64

	for i := 0; i < len(input); i += 8 {
		left += c2
		right += c1
		gammaInt = uint64(left)<<32 | uint64(right)
		for j := 0; j < 8 && i+j < len(input); j++ {
			input[i+j] ^= byte(gammaInt >> uint64(8*(7-j)))
		}
	}

	return string(input)
}

func GOSTrgzze(text string, key [8]uint32, initVector uint64) string {
	initVectorBytes := []byte{}
	for i := 0; i < 8; i++ {
		initVectorBytes = append(initVectorBytes, byte(initVector>>uint64(8*(7-i))))
	}

	input := []byte(text)

	for i := 0; i < len(input); i += 8 {
		initVectorBytes = []byte(GOSTpze(string(initVectorBytes), key))
		for j := 0; j < 8 && i+j < len(input); j++ {
			input[i+j] ^= initVectorBytes[j]
			initVectorBytes[j] = input[i+j]
		}
	}
	return string(input)
}

func GOSTrgzzd(text string, key [8]uint32, initVector uint64) string {
	initVectorBytes := []byte{}
	for i := 0; i < 8; i++ {
		initVectorBytes = append(initVectorBytes, byte(initVector>>uint64(8*(7-i))))
	}

	input := []byte(text)
	ciphertext := make([]byte, 8)

	for i := 0; i < len(input); i += 8 {
		initVectorBytes = []byte(GOSTpze(string(initVectorBytes), key))
		for j := 0; j < 8 && i+j < len(input); j++ {
			ciphertext[j] = input[i+j]
			input[i+j] ^= initVectorBytes[j]
			initVectorBytes[j] = ciphertext[j]
		}
	}
	return string(input)
}
