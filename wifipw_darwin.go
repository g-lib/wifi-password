package wifipw

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

func WIFIPassword(ssid string) (string, error) {
	output, err := exec.Command("sh", "-c",
		"security find-generic-password -D AirPort network password -wa  "+ssid).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func WIFISSID() (name string, err error) {
	var stdout bytes.Buffer
	cmdStr := "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport -I"
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	stdoutStr := stdout.String()
	if strings.Contains(stdoutStr, "AirPort: Off") {
		return "", errors.New("Wi-Fi is turned off")
	}
	ret := regexp.MustCompile(`^\s*SSID: (.+)\s*$`).FindStringSubmatch(stdout.String())
	if len(ret) < 1 {
		return "", errors.New("Could not get SSID")
	}
	return ret[1], nil
}
