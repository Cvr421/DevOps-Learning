package decrypt

func DecryptString(s string) string {
	decryptedstr := ""

	for _, c := range s {
		asciicode := int(c)

		character := string(rune(asciicode - 3))

		decryptedstr += character
	}
	return decryptedstr
}
