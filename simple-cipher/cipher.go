package cipher

import "strings"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

const plain = "abcdefghijklmnopqrstuvwxyz"

type Caesar string

func NewCaesar() Cipher {

}

func (c *Caesar) Encode(plain string) string {

	var cipher string

	for k, v := range plain {

		if v >= 65 && v <= 90 || v >= 97 && v <= 122 {
			c := strings.ToLower(string(v))
			idx := strings.Index(plain, c)
			cipher += string(plain[(idx+3)%26])
		} else {
			continue
		}
	}

	return cipher
}

func (c *Caesar) Decode(cipher string) string {

}
