package tests

import (
	CheckVMStatus "systemChecker/pkg/CheckVMStatus"
	vmStructs "systemChecker/pkg/structs"
	"testing"
)

func TestCheckVMStatus(t *testing.T) {
	// Create a test VM
	testVM := vmStructs.VM{
		Address: "google.com",
	}

	// Call the CheckVMStatus function with the test VM
	result := CheckVMStatus.CheckVMStatus(testVM)

	// Verify that the result is true
	if !result {
		t.Errorf("CheckVMStatus(%v) = %t; want true", testVM, result)
	}
}
