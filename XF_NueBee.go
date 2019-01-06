package main

import (
	"fmt"
	//"net"
	. "./xf_utility"
)

func main() {
	//init()
	/*
	AddLog(10, "test error")
	AddLog(20, "test warning")
	AddLog(30, "test info")
	AddLog(40, "test trace")
	AddLog(22, "test undefine")
	*/
	fmt.Println("----------------getMac---------------")
	nics, err := GetMac()
	if err == nil {
		for _, nic := range nics {
			fmt.Println(nic.Index, nic.Name, nic.Mac)
		}
	}
	fmt.Println("----------------getMacOne---------------")
	nics1 := []Nic{
		{1,"aa","1c:1b:0d:e1:12:eb"},
		{2,"bb","00:50:56:c0:00:01"},
		{3,"cc","00:50:56:c0:00:08"},
		{4,"dd","1b:1b:0d:e1:12:eb"},
	}
	result, err := GetMacOne(nics1)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Println("----------------bios uuid---------------")
	biosUuid, err := GetBiosUuid()
	if err == nil {
		fmt.Print(biosUuid)
	}
}