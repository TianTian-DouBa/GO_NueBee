package xf_utility

import (
    "net"
    "os/exec"
    //"fmt"
    "strings"
    "sort"
    "errors"
    "time"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "bytes"
)

var macReserved = []string {
    "00:0c:29", //VMWare
    "00:50:56", //VMWare
    "00:15:5d", //Hyper-V
}

var forfun string = "5h3u!j.x(ne=485j"
var dashu string = "cvbdg/fgk2#2"
var tyejj string = "hesg"

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

type Nic struct {
    Index int
    Name string
    Mac string
}

func getMac() ([]Nic, error) {
    itfs, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    nics := make([]Nic,0,10)
    for _, itf := range itfs {
        nic := new(Nic)
        nic.Index = itf.Index
        nic.Name = itf.Name
        nic.Mac = itf.HardwareAddr.String()
        if nic.Mac != "" {
            nics = append(nics, *nic)
        }
    }
    return nics, nil
}

func getMacOne(nics []Nic) (string, error) {
    filtedMacs := make([]Nic,0,10)
    macs := make([]string,0,10)
    if len(nics) > 0 {
        if len(nics) == 1 {
            return nics[0].Mac, nil
        } else { // len(nics) > 1
            for _, nic := range nics {
                filtered := true
                for _, check := range macReserved {
                    //fmt.Println("mac", nic.Mac)
                    //fmt.Println("check", check)
                    if strings.Index(nic.Mac,check) == 0 {
                        filtered = false
                        break
                    }
                }
                if filtered {
                    filtedMacs = append(filtedMacs, nic)
                }   
            } 
        }
    } else if len(nics) == 0 {
        return "", nil
    } else {
        return "", errors.New("len(nics) abnormal")
    }
    //fmt.Println("filtedMacs: ", filtedMacs)
    if len(filtedMacs) == 1 {
        return filtedMacs[0].Mac, nil
    } else if len(filtedMacs) > 1 {
        for _, nic :=  range filtedMacs {
            macs = append(macs, nic.Mac)
        }
    } else { //len(filtedMacs) == 0
        for _, nic1 :=  range nics {
            macs = append(macs, nic1.Mac)
        }
    }
    sort.Strings(macs)
    return macs[0], nil
}

func getBiosUuid() (string, error) {
    cmd := exec.Command("wmic","csproduct","get","UUID")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    sout := strings.Replace(string(out),"UUID                                  \r\r\n","",1)
    sout = strings.TrimSpace(sout)
    return sout, nil
}

func getCpuid() (string, error) {
    cmd := exec.Command("wmic","cpu","get","processorid")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    sout := strings.Replace(string(out),"ProcessorId       \r\r\n","",1)
    sout = strings.TrimSpace(sout)
    return sout, nil
}

func getHdid() (string, error) {
    cmd := exec.Command("wmic","diskdrive","get","serialnumber")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    sout := strings.Replace(string(out),"SerialNumber      \r\r\n","",1)
    sout = strings.TrimSpace(sout)
    return sout, nil
}

//rawId: raw string for activiation 
func rawId() (string, error) {
    nics, err := getMac()
    result := ""
    if err == nil {
        mac, err := getMacOne(nics)
        if err == nil && len(mac) == 17 {
            result += mac
        } else {
            return "", errors.New("get mjoy fail") //mac
        }
        biosUuid, err := getBiosUuid()
        if err == nil && len(biosUuid) == 36 {
            result += biosUuid
        } else {
            return "", errors.New("get bcbg fail") //bios
        }
        t := time.Now()
        result += "#dt#" + t.Format("2006-01-02 15:04:05") + "#/dt#"
        cpuid, err := getCpuid()
        if err == nil {
            result += "#cda#" + cpuid + "#/cda#"
        } else {
            result += "#cda##/cda#"
        }
        hdid, err := getHdid()
        if err == nil {
            result += "#hcd#" + hdid + "#/hcd#"
        } else {
            result += "#hcd##/hcd#"
        }
        return result, nil
    } else {
        return "", err
    }
}

func coating(rawid string) (string, error) {
    lenBraw := len(rawid)
    if lenBraw > 98 {
        braw := []byte(rawid)
        for i := 0; i < 53 ; i += 1 {
			if braw[i] == 58 {braw[i] = 120}
			if braw[i] == 45 {braw[i] = 107}
        }
        for k, v := range swd {
			braw[k], braw[v] = braw[v], braw[k]
        }
        return string(braw), nil
    } else {
        return "", errors.New("rawid invalid")
    }
}

func encrypt(raw string, key string) string {
    braw := []byte(raw)
    bkey := []byte(key)
    block, _ := aes.NewCipher(bkey)
    blockSize := block.BlockSize()
    braw = padding(braw, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, bkey[:blockSize])
    cryted := make([]byte, len(braw))
    blockMode.CryptBlocks(cryted, braw)
    return base64.StdEncoding.EncodeToString(cryted)
}

func padding(origText []byte, blocksize int) []byte {
    padlen := blocksize - len(origText)%blocksize
    padtext := bytes.Repeat([]byte{byte(padlen)}, padlen)
    return append(origText, padtext...)
}

//MakeATxt: create encrpted txt
func MakeATxt() string {
    raw, err := rawId()
    if err == nil {
        strCoated, err := coating(raw)
        if err == nil {
            smile := forfun + dashu + tyejj
            result := encrypt(strCoated, smile)
            return result
        } else {
            AddLog(10,"[fn]coating: Failed to generate txt for Activiation. err:" + err.Error())
            return ""
        }
    } else {
        AddLog(10,"[fn]rawId: Failed to generate txt for Activiation. err:" + err.Error())
        return ""
    }
}

func unPadding(paddedData []byte) []byte {
    length := len(paddedData)
    unpadlen := int(paddedData[length-1])
    return paddedData[:(length - unpadlen)]
}

func decrypt(raw string, key string) string {
    braw, _ := base64.StdEncoding.DecodeString(raw)
    bkey := []byte(key)
    block, _ := aes.NewCipher(bkey)
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, bkey[:blockSize])
    orig := make([]byte, len(braw))
    blockMode.CryptBlocks(orig, braw)
    orig = unPadding(orig)
    return string(orig)
}

func OpenATxt(raw string) (string, error) {
    key := "5h3u!j.x(ne=485jcvbdg/fgk2#2hesg"
    aTxt := decrypt(raw, key)
    lenRaw := len(aTxt)
    if lenRaw > 98 {
        braw := []byte(aTxt)
        for k, v := range swd {
			braw[k], braw[v] = braw[v], braw[k]
        }
        for i := 0; i < 53 ; i += 1 {
			if braw[i] == 120 {braw[i] = 58}
			if braw[i] == 107 {braw[i] = 45}
        }
        return string(braw), nil
    } else {
        return "", errors.New("raw invalid")
    }
}