package xf_utility

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"errors"
)

var decrypted string
var privateKey, publicKey []byte

var privateKeyS string = `
MIICXQIBAAKBgQDGqvxqlcoqLd9oHYjriUE9SNWe0YVPztEsWDrDJnd9KABFhk55
7CEVFrrKQLSAv9b5tOP+nj9Xc29P5RwPzZM9TqYt2eqLAmjZo2l9ZVE6g31RJSXD
oqVkjhmAgdjzWgmZD/ap4Wjh8Ogkir3iaZ1s6Iw0wev1SQntEzOjwqno0QIDAQAB
AoGBALPpqMmFeR2ViVoqVRKoq/IqAXrgV5AcxYUZKa4NKeynywcrR+pWEuecO/Bp
hI2MLczZKhgti0NRAJ8j5rPRAeLLTruG168oIwi5ofW/j5+g0JV3H0GLi0RT4Ycg
5/i14F57B5gJ3mveQhq+ss/k2EICW+xMxEsHi8dU6v1fBMOZAkEA9hrCi3GTp1g7
O1ULLRE2gOnnHX2vHh5IYSrKZej1lcaLCUbDo4RVKQ+CGimK69JMUyRgrqQEQtQs
dhK7LCiKhwJBAM6n8cHDrJr4qsZF2IkYLAlWnIGrt8AXnjPBwu06Lu5vA+90cF4c
okIoL7wIAdiZMRWBTOzPxWzW42/n0LERD+cCQQCs+NXviB5NKfhHlVhLjqOOK9fi
pGmmc4ZPtPGYewnRAUDLfk8W6HWqbFn25Wfco2w9q33AgUr1ZYbyXevr93qVAkBZ
8NgUN8Boli9lKcLrL90Cl3J4MS9A6EaPShY3Pypr1V9GGUxKeXXZCzpzQJzlEw88
x9CeBHNtQj0sWB0I418jAkBOiM7FGOKQV8/IKBsDkA6QTwKvPBEzuRMXlqoMgk7/
a0VpRNJ1ozJGSJqcFfchb940r3hcaDf3n9R61mUCmGh0`

var publicKeyS string = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDGqvxqlcoqLd9oHYjriUE9SNWe
0YVPztEsWDrDJnd9KABFhk557CEVFrrKQLSAv9b5tOP+nj9Xc29P5RwPzZM9TqYt
2eqLAmjZo2l9ZVE6g31RJSXDoqVkjhmAgdjzWgmZD/ap4Wjh8Ogkir3iaZ1s6Iw0
wev1SQntEzOjwqno0QIDAQAB`

func GenRsaKey(bits int) error {
	//generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("NB_private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	//generate public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block {
		Type: "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("NB_public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func genLicense(msg string) string {
	b64Msg := base64.RawStdEncoding.EncodeToString([]byte(msg))
	h := sha256.New()
	io.WriteString(h,b64Msg)
	sha256Byte := h.Sum(nil)
	sig, _ := RsaSignature(sha256Byte)
	license := base64.RawStdEncoding.EncodeToString(sig)
	return license
}

//私钥签名
func RsaSignature(message []byte) ([]byte, error) {
	rng := rand.Reader
	hashed := sha256.Sum256(message)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("invalid privateKey")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	signature, err := rsa.SignPKCS1v15(rng, private,crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

//公钥验证
func RsaSignatureVerify(message []byte, signature []byte) error {
	hashed := sha256.Sum256(message)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err := errors.New("invalid privateKey")
		return err
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}