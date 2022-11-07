package main

import "C"

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {}

var X string = getOsVersion(runtime.GOOS)

//export GetX
func GetX() *C.char {
	return C.CString(X)
}

//export SetX
func SetX(os *C.char) {
	X = C.GoString(os)
}

func getOsVersion(os string) string {
	switch os {
	case "windows":
		return getWinOsVersion()
	default:
		return os
	}
}

func getWinOsVersion() string {

	out, err := exec.Command("systeminfo").Output()
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range strings.Split(string(out), "\r\n") {
		if strings.Contains(o, "OS Name:") {
			osName := strings.SplitAfter(strings.Replace(o, "Microsoft", "", 1), "OS Name:")[1]
			return strings.TrimSpace(osName)
		}
	}

	return ""
}
