package main

import "C"
import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"sort"
	"strings"
	"time"

	"golang.org/x/sys/windows/registry"
)

var macReserved = []string{
	"00:0c:29", //VMWare
	"00:50:56", //VMWare
	"00:15:5d", //Hyper-V
}

var revId string = "Lite_0.10a"
var forfun string = "5h3u!j.x(ne=485j"
var dashu string = "cvbdg/fgk2#2"
var tyejj string = "hesg"
var rtnstr string = "MTshea8TvrR59bS55w6LCYPqvlk5oHRF"

var swd = map[int]int{
	0:  23,
	3:  12,
	5:  34,
	8:  19,
	9:  51,
	11: 25,
	13: 21,
	17: 26,
	20: 43,
	31: 47,
}

type Nic struct {
	Index int
	Name  string
	Mac   string
}

func main() {
	AddLog(20,"test log into file")
	start_s := `2018/10/31 22:28:12`
	end_s := `2018/10/31 22:58:12`
	trends_s := "SIM-001/SIN.CV, SIM-001/RAMP.CV, , , V1-COMMON/BATCH_ID.CV"
	compare := true
	startTime2_s := `2018/08/19 10:15:17`

	success := TrsPlot(start_s,end_s,trends_s, compare, startTime2_s)
	if success != true {
		fmt.Println(`failed`)
	}
	fmt.Println(`run finished`)
}

func getMac() ([]Nic, error) {
	itfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	nics := make([]Nic, 0, 10)
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
	filtedMacs := make([]Nic, 0, 10)
	macs := make([]string, 0, 10)
	if len(nics) > 0 {
		if len(nics) == 1 {
			return nics[0].Mac, nil
		} else { // len(nics) > 1
			for _, nic := range nics {
				filtered := true
				for _, check := range macReserved {
					//fmt.Println("mac", nic.Mac)
					//fmt.Println("check", check)
					if strings.Index(nic.Mac, check) == 0 {
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
		for _, nic := range filtedMacs {
			macs = append(macs, nic.Mac)
		}
	} else { //len(filtedMacs) == 0
		for _, nic1 := range nics {
			macs = append(macs, nic1.Mac)
		}
	}
	sort.Strings(macs)
	return macs[0], nil
}

func getBiosUuid() (string, error) {
	cmd := exec.Command("wmic", "csproduct", "get", "UUID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	sout := strings.Replace(string(out), "UUID                                  \r\r\n", "", 1)
	sout = strings.TrimSpace(sout)
	return sout, nil
}

func getCpuid() (string, error) {
	cmd := exec.Command("wmic", "cpu", "get", "processorid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	sout := strings.Replace(string(out), "ProcessorId       \r\r\n", "", 1)
	sout = strings.TrimSpace(sout)
	sout = strings.Replace(string(sout), "  \r\r\n", "|", -1)
	return sout, nil
}

func getHdid() (string, error) {
	cmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	sout := strings.Replace(string(out), "SerialNumber", "", 1)
	sout = strings.TrimSpace(sout)
	sout = strings.Replace(sout, "  \r\r\n", "|", -1)
	return sout, nil
}

func getOs() (string, error) {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE) // error discarded for brevity
	defer key.Close()
	productName, _, _ := key.GetStringValue("ProductName") // error discarded for brevity
	osBit := 32 << (^uint(0) >> 63)                        //返回64或32位系统
	result := fmt.Sprintln(productName, osBit)
	return result, nil
}

func getTb() (string, error) {
	path := `D:\DeltaV\DVdata\Import-Export`
	result := ""
	files, err := ioutil.ReadDir(path)
	if err != nil {
		AddLog(20, "[fn]getTb, failed to get TB info. err: "+err.Error())
		return "", err
	} else {
		for _, file := range files {
			if file.IsDir() {
				if result != "" {
					result += `|`
				}
				result += file.Name()
			}
		}
		return result, nil
	}
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
		result += "#rv#" + revId + "#/rv#"
		t := time.Now()
		result += "#dt#" + t.Format("2006-01-02 15:04:05") + "#/dt#"
		os, err := getOs()
		if err == nil {
			os = strings.TrimSpace(os)
			result += "#os#" + os + "#/os#"
		} else {
			result += "#os##/os#"
		}
		tbRev, err := getTb()
		if err == nil {
			result += "#tb#" + tbRev + "#/tb#"
		} else {
			result += "#tb##/tb#"
		}
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

func coating(rawid string) ([]byte, error) {
	lenBraw := len(rawid)
	if lenBraw > 98 {
		braw := []byte(rawid)
		for i := 0; i < 53; i += 1 {
			if braw[i] == 58 {
				braw[i] = 120
			}
			if braw[i] == 45 {
				braw[i] = 107
			}
		}
		for k, v := range swd {
			braw[k], braw[v] = braw[v], braw[k]
		}
		return braw, nil
	} else {
		return nil, errors.New("rawid invalid")
	}
}

func unPadding(paddedData []byte) []byte {
	defer func() {
		if err := recover(); err != nil {
			AddLog(10, "[fn]unPadding: Failed")
		}
	}()
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

func makeATxt() ([]byte, error) {
	raw, err := rawId()
	if err == nil {
		bytesCoated, err := coating(raw)
		if err == nil {
			result, err := RsaEncrypt(bytesCoated)
			if err == nil {
				return result, nil
			} else {
				AddLog(10, "[fn]MakeATxt: Failed to generate txt for Activiation. err:"+err.Error())
				return nil, err
			}

		} else {
			AddLog(10, "[fn]MakeATxt: Failed to generate txt for Activiation. err:"+err.Error())
			return nil, err
		}
	} else {
		AddLog(10, "[fn]MakeATxt: Failed to generate txt for Activiation. err:"+err.Error())
		return nil, err
	}
}

//WriteActiveFile: generate machine file for activiation
//export WriteActiveFile
func WriteActiveFile() (failed C._Bool) {
	path := `.\nb_active.dt`
	raw, err := makeATxt()
	if err != nil {
		AddLog(10, "[fn]WriteActiveFile fail to makeATxt. error:"+err.Error())
		failed = true
		return
	} else {
		err = WriteFileIoutil(path, raw)
		if err != nil {
			AddLog(10, "[fn]WriteActiveFile fail to write. error:"+err.Error())
			failed = true
			return
		} else {
			failed = false
			return
		}
	}
}

//export ValidKey
func ValidKey() (match C._Bool) {
	path := `.\ActiveKey.dt`
	data, err := ReadFileAll(path)
	if err == nil {
		deData := decrypt(string(data), forfun+dashu+tyejj)
		macs, _ := getMac()
		mc, _ := getMacOne(macs)
		bi, _ := getBiosUuid()
		strComp := mc + `|` + bi + `|` + rtnstr
		if deData == strComp {
			match = true
			//failed = false
			return
		} else {
			match = false
			//failed = false
			return
		}
	} else {
		AddLog(10, "[fn]ValidKey fail to read file. error:"+err.Error())
		match = false
		//failed = true
		return
	}
}

//export TrsPlot
func TrsPlot(start_s string, end_s string, trends_s string, compare bool, startTime2_s string) (success C._Bool) {
	AddLog(40, "[fn]TrsPlot.start_s:"+start_s)
	AddLog(40, "[fn]TrsPlot.end_s:"+end_s)
	var err error
	if compare == false {
		err = runChs(start_s, end_s, trends_s)
		if err !=nil {
			success = false
			AddLog(10, "[fn]TrsPlot.error:"+err.Error())
		} else {
			success = true
		}
	} else if compare == true {
		err = compChs(start_s, end_s, trends_s, startTime2_s)
		if err !=nil {
			success = false
			AddLog(10, "[fn]TrsPlot.error:"+err.Error())
		} else {
			success = true
		}
	}
	return
}
