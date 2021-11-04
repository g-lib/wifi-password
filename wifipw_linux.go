package wifipw

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func WIFIPassword(ssid string) (password string, err error) {
	paths := []string{
		fmt.Sprintf("/etc/NetworkManager/system-connections/%s", ssid),
		fmt.Sprintf("/etc/NetworkManager/system-connections/%s.nmconnection", ssid),
	}
	outputStr := ""
	for _, p := range paths {
		output, err := exec.Command("sudo", "cat", p).CombinedOutput()
		if err == nil {
			outputStr = string(output)
			break
		}
	}
	if outputStr == "" {
		return "", errors.New("could not get password")
	}
	ret := regexp.MustCompile(`(?:psk|password)=(.+)`).FindStringSubmatch(outputStr)
	if len(ret) < 1 {
		return "", errors.New("could not get SSID")
	}
	return strings.TrimSpace(ret[1]), nil
}

func WIFISSID() (name string, err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", "iwgetid -r")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", errors.New(stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}
