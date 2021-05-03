package keygen

import "fmt"

type KeyGenerator struct {
}

func New() KeyGenerator {
	return KeyGenerator{}
}
func (k KeyGenerator) GenerateKey(url string) string {
	return fmt.Sprintf("%d", len(url))
}
