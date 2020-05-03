package main

import "testing"

func TestGetMd5URL(t *testing.T) {
	rawURL := "https://adjust.com"

	expected := "4c1353caaba3c2329f411cde6fa412a6"

	if got, _ := GetMd5URL(rawURL); got != expected {
		t.Errorf("%s = %s; want %s", rawURL, got, expected)
	}

	if _, err := GetMd5URL("incorrectulr"); nil == err {
		t.Errorf(err.Error())
	}
}

