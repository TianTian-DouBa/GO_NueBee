package main

import (
	"io/ioutil"
	//"os"
)

/* FileMode 八进制 rwx, User Group Other
0644: user(rw) group(r) other(r)
0777: user(rwx) group(rwx) other(rwx)
*/

//ReadFileAll: read all the content of the file
func ReadFileAll(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		errorStr := "[fn]ReadFileAll: err.Error()"
		AddLog(20, errorStr)
		return nil, err
	} else {
		return data, nil
	}
}

func WriteFileIoutil(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	} else {
		return nil
	}
}
