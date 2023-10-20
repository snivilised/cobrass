package clif_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClif(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clif Suite")
}
