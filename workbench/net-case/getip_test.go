package main

import "testing"

func TestGetIp(t *testing.T) {
	ips, err := Ips()
	if err != nil {
		t.Fatalf("Got err: %s", err)
	}
	if len(ips) == 0 {
		t.Fatalf("ips length 0")
	}
	for _, ip := range ips {
		t.Logf("%s\n", ip)
	}
}
