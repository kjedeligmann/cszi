package main

import (
	"bufio"
	"log"
	"os"
)

func LFSR(start uint64, polynomial uint64, filename string) {
	read, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	write, err := os.Create(filename + ".lfsr")
	if err != nil {
		log.Fatal(err)
	}
	defer write.Close()

	lfsr := start
	var order uint8

	for taps := polynomial; taps != 0; taps >>= 1 {
		order++
	}
	polynomial ^= (1 << (order - 1))

	reader := bufio.NewReader(read)
	writer := bufio.NewWriter(write)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		g := LFSRbyte(&lfsr, &polynomial, &order)
		writer.WriteByte(b ^ g)
	}

	writer.Flush()
}

func LFSRbyte(register *uint64, polynomial *uint64, order *uint8) (result byte) {
	for i := 7; i >= 0; i-- {
		result |= byte((*register & 1) << i)
		var bit uint64
		bits := *register & *polynomial
		for ; bits != 0; bits >>= 1 {
			bit ^= bits
		}
		bit &= 1
		*register = (*register >> 1) | (bit << (*order - 1))
	}
	return
}

func main() {
	var start, polynomial uint64 = 0xACE1, 0b1000000000101101
	LFSR(start, polynomial, "hello.txt")
	LFSR(start, polynomial, "hello.txt.lfsr")
}
