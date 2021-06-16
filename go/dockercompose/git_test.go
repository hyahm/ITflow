package dockercompose

import (
	"bytes"
	"testing"

	"github.com/go-git/go-git/v5"
	sshgit "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/hyahm/golog"
	"golang.org/x/crypto/ssh"
)

func TestGit(t *testing.T) {
	defer golog.Sync()
	pri := `
	-----BEGIN OPENSSH PRIVATE KEY-----
	  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
	-----END OPENSSH PRIVATE KEY-----
	
	`

	temp := []byte(pri)
	golog.Info(temp)
	temp = bytes.ReplaceAll(temp, []byte{9}, []byte{})
	signer, err := ssh.ParsePrivateKey(temp)
	if err != nil {
		t.Log(err)
		return
	}

	auth := &sshgit.PublicKeys{
		User:   "git",
		Signer: signer,
		HostKeyCallbackHelper: sshgit.HostKeyCallbackHelper{
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}
	_, err = git.PlainClone("aaaaaaaaaaaaaaaaa", false, &git.CloneOptions{
		URL:  "git@47.107.77.123:shandiao/xilin-backend.git",
		Auth: auth,
	})
	if err != nil {
		if err == git.ErrRepositoryAlreadyExists {
			golog.Info("11111111111")
		}
		t.Log(err)
		return
	}
	t.Log("commmmmm")
}
