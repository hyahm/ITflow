// include rsa md5 sha1 encrypt
package gaencrypt

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

func PwdEncrypt(str string, salt string) (x string) {
	w := md5.New()
	io.WriteString(w, str+salt)
	m := fmt.Sprintf("%x", w.Sum(nil))

	s := sha1.New()
	io.WriteString(s, m)
	x = fmt.Sprintf("%x", s.Sum(nil))
	return

}

func Md5(s string, salt string) string {
	m := sha1.New()
	io.WriteString(m, s+salt)
	return string(m.Sum(nil))
}

func Sha1(s string, salt string) string {
	sha := sha1.New()
	io.WriteString(sha, s+salt)
	return string(sha.Sum(nil))
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

func RsaDecryptByte(pwd string, privatekey string, b64 bool) ([]byte, error) {
	var mm []byte
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
