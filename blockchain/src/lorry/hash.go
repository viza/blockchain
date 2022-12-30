package lorry

import (
	"crypto/sha1"
	"io"

	"github.com/gookit/slog"
)

type hash struct {
	hash string
}

// ToSHA1 implements Hasher
func (*hash) ToSHA1(message string) string {

	h := sha1.New()
	slog.Warn("TODO: Consider other sha algorithms in future")
	io.WriteString(h, message)
	slog.Infof("SHA1: % x\n", h.Sum(nil))
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
