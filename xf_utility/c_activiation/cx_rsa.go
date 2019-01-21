package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var decrypted string
var publicKey []byte

var publicKeyS string = `-----BEGIN RSA PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAzTn+tOhb7C/qTiSeWddB
yROGiwkLIXifW285sM+y0y3FoEE5lRG7ZXBfNsqdDrKx0wJZS3XsqxFPMJdemTPA
oiXwDCfIHfe/zNgQwEDNhcXrwwvA3P+9GyH/aw8wJy5r/qez8lOCZRZJ4nvi29hp
a22eDHct4i1wIRJqW3k1nwdm1RyK7uvYiUycDD+T4R/HbXgpmB980b/JIquG6aS1
wxMLckmz0ITARAI1JPXUP8JNOfKh1yFXtohN3D2jkxymaqkMmNx17do99VdJ5xBz
aep2UFIhVaWM/UNl3zPJNPIcqVHi2Y+KOV4uorr/dPDW89cvZMWYryXU9bZtM4Lo
hIc2OLn5dl93VCQOBsN6gx6PZcclJN8MRVch3s0u6IuXair5PlfCVqf1nQx9JeZ2
DGNCsyHG7juk7iZEJpRqr5Iu9beol7LLtpojrxGgTLPhMkMsozXkSHz2LFEbKN9H
KXn88TeZpjVyOtKHj9k096LIlWpUJ7ls/nmV3BQIYCAug1+vyplwFJroDkvspE0g
gveohkK+2xxQAeSrw6qPfODrqz5v/+tpPSBnZz4oQYR8Vh0dbX3/KfImGKgA5GKN
RSeeQXvPd9hA6Y5DRDmd3Id4xgrV6gzaRB1zRdAdZAe9lIZM5tWfoWvBRIwC89JU
RNqXG10JfQUyhr3+iQpCbBsCAwEAAQ==
-----END RSA PUBLIC KEY-----

`

//公钥加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyS))
	if block == nil {
		return nil, errors.New("invalid publicKey")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
