package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func randStr(size int) (string, error) {
	sl := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, sl); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(sl), nil
}

func main() {
	println(randStr(16))
}
