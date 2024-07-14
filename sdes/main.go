package main

import "fmt"

func main() {
	fmt.Println("Testing SDES encryption and decryption")
	cryptograms := []struct {
		Plaintext string
		Key       bit10
	}{
		{"Keep your friends close, but your enemies closer.", 296},
		{"Injustice anywhere is a threat to justice everywhere.", 618},
		{"Life is what happens when you're busy making other plans.", 254},
	}

	for _, c := range cryptograms {
		ciphertext := SDESe(c.Plaintext, c.Key)
		decrypted := SDESd(ciphertext, c.Key)

		fmt.Println("Plaintext: ", c.Plaintext)
		fmt.Println("Ciphertext:", ciphertext)
		fmt.Println("Decrypted: ", decrypted)
		fmt.Printf("Plaintext (hex):  %x\n", c.Plaintext)
		fmt.Printf("Ciphertext (hex): %x\n", ciphertext)
		fmt.Printf("Decrypted (hex):  %x\n", decrypted)
	}
}
