package main

import (
	"os/exec"
	"log"
)

var CHS_PATH string = `C:\DeltaV\Bin\CHS.exe`

//runChs: plot the standard PHV trends
func runChs(start_s string,end_s string,trends_s string) error {
	cmd := exec.Command(CHS_PATH,`/new`,`EChart`,`/starttime`,start_s,`/endtime`,end_s,`/trend`, trends_s)
	//cmd := exec.Command(CHS_PATH,`/new`,` EChart`,` /starttime `,start_s)
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	return nil
}

//compChs: plot the PHV trends with comparation
func compChs(start_s string,end_s string,trends_s string, startTime2_s string) error {
	cmd := exec.Command(CHS_PATH,`/new`,`EChart`,`/starttime`,start_s,`/endtime`,end_s,`/trend`, trends_s,`/starttime2`, startTime2_s)
	//cmd := exec.Command(CHS_PATH,`/new`,` EChart`,` /starttime `,start_s)
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	return nil
}