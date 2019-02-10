package main

import (
	"os/exec"
	"log"
)

var CHS_PATH string = `C:\DeltaV\Bin\CHS.exe`

func runChs(start_s string,end_s string,trends_s string) error {
	cmd := exec.Command(CHS_PATH,`/new`,`EChart`,`/starttime`,start_s,`/endtime`,end_s,`/trend`, trends_s)
	//cmd := exec.Command(CHS_PATH,`/new`,` EChart`,` /starttime `,start_s)
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	return nil
}