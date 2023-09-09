package gola_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGola(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gola Suite")
}
