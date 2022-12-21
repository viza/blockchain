package lorry

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type hash struct {
	hash string
}

// ToSHA1 implements Hasher
func (*hash) ToSHA1(message string) string {

	h := sha1.New()
	io.WriteString(h, message)
	fmt.Printf("SHA1: % x\n", h.Sum(nil))
	return string(h.Sum(nil))
}

type Hasher interface {
	ToSHA1(message string) string
}

func Hash() Hasher {
	return &hash{
		hash: "",
	}
}
