package emitter_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEmitter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Emitter Suite")
}
