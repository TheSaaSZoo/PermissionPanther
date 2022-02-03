package main

import "testing"

func HandleTestError(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}
