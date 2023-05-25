package main

import (
	"fmt"
	"go-rc6/rc6"
)

func main() {
	your_encryption_key := "ashiap"
	your_plaintext_data_max_16_character := "achmadrizkigante"

	key := []byte(your_encryption_key)
	cipherRC6 := rc6.NewCipher(key)

	fmt.Println("ENCRYPTION KEY: " + your_encryption_key)
	fmt.Println("YOUR PLAIN TEXT DATA: " + your_plaintext_data_max_16_character)

	// encription
	if len(your_plaintext_data_max_16_character) > 16 {
		fmt.Println("plain text cannot max 16")
	} else {
		plaintext := []byte(your_plaintext_data_max_16_character)
		ciphertext := make([]byte, len(your_plaintext_data_max_16_character))

		cipherRC6.Encrypt(ciphertext, plaintext)

		fmt.Println("")
		fmt.Println("ENCRYPTION STRING: ")
		fmt.Println(string(ciphertext))
		fmt.Println("ENCRYPTION BYTE: ")
		fmt.Println(ciphertext)

		decryptedPlaintext := make([]byte, len(your_plaintext_data_max_16_character))
		cipherRC6.Decrypt(decryptedPlaintext, ciphertext)

		fmt.Println("")
		fmt.Println("DECRYPTED STRING: ")
		fmt.Println(string(decryptedPlaintext))
		fmt.Println("DECRYPTED BYTE: ")
		fmt.Println(decryptedPlaintext)
	}
}
