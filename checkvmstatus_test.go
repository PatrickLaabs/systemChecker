package main

import (
	"testing"
	"time"
)

func TestCheckVMStatus(t *testing.T) {
	// Test for an available VM
	available := CheckVMStatus("192.168.2.180", time.Second)
	if !available {
		t.Errorf("checkVMStatus failed for available VM")
	}

	// Test for an unavailable VM
	unavailable := CheckVMStatus("192.168.2.221", time.Second)
	if unavailable {
		t.Errorf("checkVMStatus failed for unavailable VM")
	}
}
