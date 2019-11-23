package main

import (
	"fmt"
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

	Context("Cloning a repo", func() {
		It("with https url", func() {
			_, err = git.Clone("https://github.com/libgit2/git2go.git", tmpRepoDir, &git.CloneOptions{})
			Expect(err).ToNot(HaveOccurred())
		})

		It("with ssh", func() {
			publicKeyPath := ""
			privateKeyPath := ""
			passphrase := ""
			options := &git.CloneOptions{
				FetchOptions: &git.FetchOptions{
					RemoteCallbacks: git.RemoteCallbacks{
						CredentialsCallback: func(url string, username_from_url string, allowed_types git.CredType) (git.ErrorCode, *git.Cred) {
							i, c := git.NewCredSshKey(username_from_url, publicKeyPath, privateKeyPath, passphrase)
							return git.MakeGitError2(i).(*git.GitError).Code, &c
						},
						CertificateCheckCallback: func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
							if valid {
								return git.ErrOk
							}
							return git.ErrOk
						},
					},
				},
			}
			_, err = git.Clone("git@github.com:libgit2/git2go.git", tmpRepoDir, options)
			Expect(err).ToNot(HaveOccurred())
		})

	})

	FContext("rev-list", func() {
		It("lists commits", func() {
			repo, err := git.OpenRepository("/Users/taylor/workspace/concourse")
			Expect(err).ToNot(HaveOccurred())

			oid, err := git.NewOid("eb4c434bba9ba2e256241e9fe4797ff46226f2e5")
			Expect(err).ToNot(HaveOccurred())
			commit, err := repo.LookupCommit(oid)
			Expect(err).ToNot(HaveOccurred())
			tree, err := commit.Tree()
			Expect(err).ToNot(HaveOccurred())
			fmt.Println("tree:", tree.EntryCount())
		})
	})
})
