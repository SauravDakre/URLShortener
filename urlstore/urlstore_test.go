package urlstore

import (
	"github.com/SauravDakre/URLShortener/internal/mock/keygen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URLStore", func() {
	var u URLStore

	BeforeEach(func() {
		u = NewURLStore()
	})

	It("should store key, value pair in map", func() {

		key := "a"
		val := "abc.com"
		u.set(key, val)
		Expect(len(u.urls)).To(Equal(1))
		Expect(u.urls[key]).To(Equal(val))
	})

	It("should store key, value pair in map", func() {
		key := "g"
		val := "google.com"
		u.set(key, val)
		url := u.Get(key)
		Expect(val).To(Equal(url))
	})

	It("should put value in url store", func() {
		url := "google.com"
		k := keygen.New()
		key := u.Put(url, k)
		val := u.Get(key)
		Expect(len(u.urls)).To(Equal(1))
		Expect(val).To(Equal(url))
	})
})
