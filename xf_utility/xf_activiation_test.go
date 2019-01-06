package xf_utility

import (
	"testing"
)

func TestGetMacOne(t *testing.T) {
	nics1 := []Nic{
		{1,"aa","1c:1b:0d:e1:12:eb"},
		{2,"bb","00:50:56:c0:00:01"},
		{3,"cc","00:50:56:c0:00:08"},
		{4,"dd","1b:1b:0d:e1:12:eb"},
	}
	result, err := GetMacOne(nics1)
	if (err == nil) && (result == "1b:1b:0d:e1:12:eb") {
		t.Log("pass")
	} else {
		t.Error("failed")
	}

	nics2 := []Nic{
		{1,"aa","00:50:56:e1:12:eb"},
		{2,"bb","00:50:56:c0:00:01"},
		{4,"dd","00:15:5d:e1:12:eb"},
		{3,"cc","00:50:56:c0:00:08"},
	}
	result, err = GetMacOne(nics2)
	if (err == nil) && (result == "00:15:5d:e1:12:eb") {
		t.Log("pass")
	} else {
		t.Error("failed")
	}

	nics3 := []Nic{}
	result, err = GetMacOne(nics3)
	if (err == nil) && (result == "") {
		t.Log("pass")
	} else {
		t.Error("failed: ", err)
	}

	nics4 := []Nic{
		{1,"aa","00:50:56:e1:12:eb"},
		{2,"bb","45:50:56:c0:00:01"},
		{3,"dd","00:15:5d:e1:12:eb"},
		{4,"cc","00:50:56:c0:00:08"},
		{5,"aa","19:50:56:e1:12:eb"},
		{6,"bb","27:50:56:c0:00:01"},
		{7,"dd","36:15:5d:e1:12:eb"},
		{8,"cc","28:50:56:c0:00:08"},
		{9,"aa","19:50:56:e1:12:aa"},
		{10,"bb","27:50:56:c0:00:aa"},
		{11,"dd","36:15:5d:e1:12:aa"},
		{12,"cc","28:50:56:c0:00:aa"},
	}
	result, err = GetMacOne(nics4)
	if (err == nil) && (result == "19:50:56:e1:12:aa") {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}