package CheckVMStatus

import (
	"os/exec"
	vmStructs "systemChecker/pkg/structs"
)

func CheckVMStatus(vm vmStructs.VM) bool {
	cmd := exec.Command("ping", "-c", "1", vm.Address)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
