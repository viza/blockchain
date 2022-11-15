package main

import (
	"fmt"
)

func main() {
	// Verification service: https://cryptii.com/pipes/vigenere-cipher
	testStr := "Russia is a terrorist country. Punin hujlo"
	key := "TestKey"

	fmt.Println("String to enrypt:", testStr)
	encodedStr := encode(key, testStr)
	fmt.Println("EncodedStr:", encodedStr)

	decodedStr := decode(key, encodedStr)
	fmt.Println("DecodedStr:", decodedStr)

}

func encode(key, msg string) string {
	skey, smsg := sanitize(key), sanitize(msg)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, encodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func sanitize(str string) string {
	out := []rune{}
	for _, v := range str {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}
	return string(out)
}

func encodePair(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func decode(key, msg string) string {
	skey, smsg := sanitize(key), sanitize(msg)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, decodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func decodePair(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}
