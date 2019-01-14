package main

import (
	"fmt"
	. "../xf_utility"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"regexp"
)

var privateKey []byte
var privateKeyS string = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`
var smile string = "Lite_0.10a5h3u!j.x(ne=485jcvbdg/fgk2#2"
var swd = map[int]int {
    0 : 23,
    3 : 12,
    5 : 34,
    8 : 19,
    9 : 51,
    11 : 25,
    13 : 21,
    17 : 26,
    20 : 43,
    31 : 47,
}
var rtnstr string = "MTshea8TvrR59bS55w6LCYPqvlk5oHRF"

func main() {
	path := `.\nb_active.dt`
	data, err := ReadFileAll(path)
	if err == nil {
		deData, err := RsaDecrypt(data)
		if err == nil {
			result, err := decoating(deData)
			if err == nil {
				fmt.Println(string(result))
				aHost := extractInfo(result)
				fmt.Println("-----------------")
				fmt.Println(aHost.Mac)
				fmt.Println(aHost.BiosUuid)
				fmt.Println(aHost.NbRev)

			} else {
				AddLog(10,"[fn]main: Failed to generate key:" + err.Error())
			}
			
		}
	}
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

//去掉混淆
func decoating (coated []byte) ([]byte, error) {
	lens := len(coated)
	if lens > 98 {
		for k, v := range swd {
			coated[k], coated[v] = coated[v], coated[k]
		}
		for i := 0; i < 53 ; i += 1 {
			if coated[i] == 120 {coated[i] = 58}
			if coated[i] == 107 {coated[i] = 45}
		}
		return coated, nil
	} else {
		return nil, errors.New("coated input invalid")
	}
}

type Activiator struct {
	Mac string
	BiosUuid string
	NbRev string
	ExpDate string
	OpSys string
	TbRev string
	Cpu	string
	Disk string
}

//提取信息
func extractInfo(bInfo []byte) (*Activiator) {
	AHost := new(Activiator)
	AHost.Mac = string(bInfo[:17])
	AHost.BiosUuid = string(bInfo[17:53])
	strInfo := string(bInfo)
	re := regexp.MustCompile(`#rv#(.*)#/rv#`)
	foundStr := re.FindString(strInfo)
	AHost.NbRev := foundStr[4:-3]
	fmt.Println(foundStr)

	return AHost
}

//MakeKey: generate key TBD,生成的Key文件以此对称加密
/*
func MakeKey() []byte {
    raw, err := rawId()
    if err == nil {
        strCoated, err := coating(raw)
        if err == nil {
            result := []byte(encrypt(strCoated, smile))
            return result
        } else {
            AddLog(10,"[fn]coating: Failed to generate txt for Activiation. err:" + err.Error())
            return nil
        }
    } else {
        AddLog(10,"[fn]rawId: Failed to generate txt for Activiation. err:" + err.Error())
        return nil
    }
}
*/