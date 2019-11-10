package main

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	git "gopkg.in/libgit2/git2go.v27"
)

var _ = Describe("Cloning a git repo", func() {
	var (
		tmpRepoDir string
		err        error
	)

	BeforeEach(func() {
		tmpRepoDir, err = ioutil.TempDir("/tmp", "git2go-test")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		err = os.RemoveAll(tmpRepoDir)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("Cloning a http url", func() {
		It("with https url", func() {
			_, err = git.Clone("https://github.com/libgit2/git2go.git", tmpRepoDir, &git.CloneOptions{})
			Expect(err).ToNot(HaveOccurred())
		})

	})
})
