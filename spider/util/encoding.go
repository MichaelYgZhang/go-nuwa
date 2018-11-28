package util

import (
	"bufio"
	"log"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("determineEncoding error: %v", err)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}
