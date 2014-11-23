package util

import (
	"crypto/rand"
	"encoding/hex"
)

const XML_DECLARATION = `<?xml version="1.0" encoding="UTF-8"?>`

func Id() (string, error) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	return hex.EncodeToString(b), err
}
