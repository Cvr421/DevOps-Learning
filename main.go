package main

import (
	"fmt"

	"cryptit/decrypt"
	"cryptit/encrypt"
)

func main() {
	encryptStr := encrypt.EncryptString("Hello World")

	fmt.Println(encryptStr)
	fmt.Println(decrypt.DecryptString(encryptStr))
}
