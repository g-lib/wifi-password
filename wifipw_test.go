package wifipw_test

import (
	"testing"

	wifipw "github.com/g-lib/wifi-password"
)

func TestWIFISSID(t *testing.T) {
	t.Log(wifipw.WIFISSID())
}

func TestWIFIPassword(t *testing.T) {
	ssid, _ := wifipw.WIFISSID()
	t.Log(wifipw.WIFIPassword(ssid))
}

func TestGetWIFIPassword(t *testing.T) {
	t.Log(wifipw.GetWIFIPassword())
}
