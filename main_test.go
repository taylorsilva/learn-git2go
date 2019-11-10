package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelloWorld(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Learning git2go")
}
