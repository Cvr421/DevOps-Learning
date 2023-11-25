package encrypt

func EncryptString(s string) string {
	encryptedstr := ""

	for _, c := range s {
		asciicode := int(c)

		character := string(rune(asciicode + 3))

		encryptedstr += character
	}
	return encryptedstr
}
