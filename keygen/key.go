package keygen

import "crypto/sha1"

func GenerateKey(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return string(bs[:8])
}
