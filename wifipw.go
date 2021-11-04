package wifipw

func GetWIFIPassword() (ssid, password string, err error) {
	if s, err := WIFISSID(); err != nil {
		return "", "", err
	} else {
		ssid = s
	}
	if pw, err := WIFIPassword(ssid); err != nil {
		return ssid, "", err
	} else {
		password = pw
	}
	return
}
