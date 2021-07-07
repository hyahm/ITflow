package dockercompose

// caddy 配置文件的封装
import (
	"bytes"
	"errors"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	sshgit "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/hyahm/golog"
	"golang.org/x/crypto/ssh"
)

type Git struct {
	Url      string
	Path     string
	User     string // 认证账号
	Password string // 认证密码
	AuthType int    // 认证类型, 1: 密码认证， 2： 密钥认证
	Auth     transport.AuthMethod
	Pri      string
}

func isPrivateKey(pass string) bool {
	if len(pass) > 1000 && strings.HasPrefix(pass, "-----") {
		return true
	}
	return false
}

func (g *Git) getAuth() error {
	switch g.AuthType {
	case 1:
		g.Auth = &http.BasicAuth{
			Username: g.User,
			Password: g.Password,
		}
		return nil
	case 2:
		var signer ssh.Signer
		var err error
		// 去掉可能出现的无效 tab, 换行和空格， 官方已经去掉了
		temp := []byte(g.Pri)
		temp = bytes.ReplaceAll(temp, []byte{9}, []byte{})
		if g.Password == "" {
			signer, err = ssh.ParsePrivateKey(temp)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(temp, []byte(g.Password))
		}

		if err != nil {
			golog.Error(err)
			return err
		}
		if g.User == "" {
			g.User = "git"
		}
		g.Auth = &sshgit.PublicKeys{
			User:   g.User,
			Signer: signer,
			HostKeyCallbackHelper: sshgit.HostKeyCallbackHelper{
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
		}
		return nil
	default:
		return errors.New("not support")
	}

}

// 拉取更新数据
func (g *Git) GitPull() error {
	if g.Path == "" {
		return git.ErrRemoteNotFound
	}
	repo, err := git.PlainOpen(g.Path)
	if err != nil {
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}
	pull := &git.PullOptions{
		Auth:  g.Auth,
		Force: true,
	}
	if err := g.getAuth(); err != nil {
		return err
	}
	pull.Auth = g.Auth
	return wt.Pull(pull)
}

func (g *Git) GitClone() error {
	if g.Url == "" {
		return git.ErrMissingURL
	}
	clone := &git.CloneOptions{
		URL: g.Url,
	}
	if err := g.getAuth(); err != nil {
		return errors.New("auth failed")
	}
	clone.Auth = g.Auth
	golog.Info(g.Path)
	_, err := git.PlainClone(g.Path, false, clone)
	return err
}
