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

var privateKeyS string = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----

`

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
	block, _ := pem.Decode([]byte(privateKeyS))
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

//公钥验签
func RsaSignatureVerify(message []byte, signature []byte) error {
	hashed := sha256.Sum256(message)
	block, _ := pem.Decode([]byte(publicKeyS))
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

// 私钥解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
    block, _ := pem.Decode([]byte(privateKeyS))
    if block == nil {
        return nil, errors.New("invalid privateKey")
    }
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

	