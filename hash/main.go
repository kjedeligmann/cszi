package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

/*
Write a CLI tool that implements a hashing algorithm in your preferred programming language.

| Algorithm | Message source | Input |    Output    |
|-----------|----------------|-------|--------------|
|   SHA-1   |      File      |  Hex  | Hash(in hex) |
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read all hex data from the file
	hexData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Delete the file terminator character
	hexData = hexData[:(len(hexData) - 1)]

	// Decode the hex data into bytes
	rawData := make([]byte, hex.DecodedLen(len(hexData)))
	_, err = hex.Decode(rawData, hexData)
	if err != nil {
		log.Fatalf("Error decoding hex data: %v", err)
	}

	// Compute SHA-1 hash
	hash := sha1.Sum(rawData)

	// Print the hash as a hex string
	hashHex := hex.EncodeToString(hash[:])
	fmt.Printf("SHA-1 Hash of the data: %s\n", hashHex)
}
