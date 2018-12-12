package util

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Sha1Encryption(in string) (out string, err error) {
	hash := sha1.New()
	if _, err = io.WriteString(hash, in); err != nil {
		return
	}
	out = fmt.Sprintf("%x",hash.Sum(nil))
	return
}