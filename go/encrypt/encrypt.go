// include rsa md5 sha1 encrypt
package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

func PwdEncrypt(str string, salt string) string {
	return Sha1(Md5(str + salt))

}

func Md5(s string) string {
	m := md5.New()
	io.WriteString(m, s)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func Sha1(s string) string {
	sha := sha1.New()
	io.WriteString(sha, s)
	return fmt.Sprintf("%x", sha.Sum(nil))
}

//
func Token(nickname string, salt string) string {

	str := rangeTime() + nickname + salt

	return PwdEncrypt(str, salt)

}

func rangeTime() string {
	t := time.Now().UnixNano()
	return strconv.FormatInt(t, 10)
}

func RsaDecrypt(pwd string, privatekey string, b64 bool) ([]byte, error) {
	if privatekey == "" {
		return nil, errors.New("私钥不能为空")
	}
	mm := []byte(pwd)
	if b64 {
		mm, _ = base64.StdEncoding.DecodeString(pwd)
	}

	block, _ := pem.Decode([]byte(privatekey))
	if block == nil {
		return nil, errors.New("private key error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, mm)
}
