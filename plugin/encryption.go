package plugin

import "strings"

type Encryption interface {
	Encrypt(input string) string
	Decrypt(input string) string
}

type Caesar struct {
	encryptionIsEnabled bool
	offset rune
}

func NewCaesar(enableEncyption bool, key int) Caesar {
	return Caesar{
		encryptionIsEnabled: enableEncyption,
		offset: rune(key),
	}
}

func (enc Caesar) Encrypt(input string) string {
	if enc.encryptionIsEnabled {
		return strings.Map(enc.encryptChar, input)
	} 
	return input
}

func (enc Caesar) encryptChar(character rune) rune {
	switch {
	case character >= 'A' && character <= 'Z':
		return 'A' + (character - 'A' + enc.offset) % 26
	case character >= 'a' && character <= 'z':
		return 'a' + (character - 'a' + enc.offset) % 26
	}
	return character
}

func (enc Caesar) Decrypt(input string) string {
	if enc.encryptionIsEnabled {
		return strings.Map(enc.decryptChar, input)
	}
	return input
}

func (enc Caesar) decryptChar(character rune) rune {
	switch {
	case character >= 'A' && character <= 'Z':
		return 'Z' - ('Z' - character + enc.offset) % 26
	case character >= 'a' && character <= 'z':
		return 'z' - ('z' - character + enc.offset) % 26
	}
	return character
}
