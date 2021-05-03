package urlstore

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUrlstore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Urlstore Suite")
}
