package wifipw

import (
	"errors"
	"os/exec"
	"regexp"
)

func WIFIPassword(ssid string) (string, error) {
	output, err := exec.Command("cmd", "/C", "netsh", "wlan",
		"show", "profile", "name="+ssid, "key=clear").CombinedOutput()
	if err != nil {
		return "", err
	}
	ret := regexp.MustCompile(`^\s*Key Content\s*: (.+)\s*$`).FindStringSubmatch(string(output))
	if len(ret) < 1 {
		return "", errors.New("Could not get password")
	}
	return ret[1], nil
}

func WIFISSID() (string, error) {
	cmd, err := exec.Command("cmd", "/C", "netsh", "wlan", "show", "interface").CombinedOutput()
	if err != nil {
		return "", err
	}
	ret := regexp.MustCompile(`^\s*SSID: (.+)\s*$`).FindStringSubmatch(string(cmd))
	if len(ret) < 1 {
		return "", errors.New("Could not get SSID")
	}
	return ret[1], nil
}
