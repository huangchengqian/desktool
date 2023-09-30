package util

import (
	"os"
)


func Convert2Byte(path string) []byte {
	srcByte, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return srcByte
}