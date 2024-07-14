package main

import (
	"testing"
)

func benchmarkCaesar(b *testing.B, plaintext string, key int) {
	for i := 0; i < b.N; i++ {
		caesarUkr(&plaintext, key)
	}
}

func BenchmarkSimple(b *testing.B) {
	benchmarkCaesar(b, "Щось", 3)
}

func BenchmarkHarder(b *testing.B) {
	benchmarkCaesar(b, "Украй простий приклад симетричного шифрування — шифр підстановки.", 3)
}

func BenchmarkHardest(b *testing.B) {
	benchmarkCaesar(b, "Украй простий приклад симетричного шифрування — шифр підстановки.Украй простий приклад симетричного шифрування — шифр підстановки.Украй простий приклад симетричного шифрування — шифр підстановки.Украй простий приклад симетричного шифрування — шифр підстановки.Украй простий приклад симетричного шифрування — шифр підстановки.Украй простий приклад симетричного шифрування — шифр підстановки.", 3)
}
