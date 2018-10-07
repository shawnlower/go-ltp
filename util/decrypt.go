package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	var filename string
	var message []byte

	if len(os.Args) != 3 {
		fmt.Println("Usage: $0 key filename")
		fmt.Println("Data will be written to stdout.")
		os.Exit(1)
	}

	filename = os.Args[2]
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(fd)
	message = make([]byte, buf.Len())
	copy(message, buf.Bytes())

	key, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Println("Invalid hex digits: ", key)
		os.Exit(1)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := message[:gcm.NonceSize()]

	os.Stderr.WriteString(fmt.Sprintf("Using %d-bit key: %x\n",
		len(key)*8, key))
	os.Stderr.WriteString(fmt.Sprintf("Using %d-bit nonce: %x\n",
		len(nonce)*8, nonce))
	// os.Stderr.WriteString(fmt.Sprintf("Ciphertext (%d bytes): %x",
	//  len(message), message))

	plaintext, err := gcm.Open(nil, nonce, message[len(nonce):], nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", plaintext)

}
