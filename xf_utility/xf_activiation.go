package xf_utility

import (
    "net"
    "os/exec"
    //"fmt"
    "strings"
    "sort"
    "errors"
)

var macReserved = []string {
    "00:0c:29", //VMWare
    "00:50:56", //VMWare
    "00:15:5d", //Hyper-V
}

type Nic struct {
    Index int
    Name string
    Mac string
}

func GetMac() ([]Nic, error) {
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

func GetMacOne(nics []Nic) (string, error) {
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

func GetBiosUuid() (string, error) {
    cmd := exec.Command("wmic","csproduct","get","UUID")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    sout := strings.Replace(string(out),"UUID                                  \r\r\n","",1)
    return sout, nil
}

// func getVmUuid() (string, error) {

// }
