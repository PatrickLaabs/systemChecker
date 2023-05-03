package CheckVMStatus

import (
	"fmt"
	"os/exec"
	vmStructs "systemChecker/pkg/structs"
)

type VM struct {
	Name         string
	Address      string
	Status       bool
	Dependencies []string
}

func CheckVMStatus(vm vmStructs.VM) bool {
	cmd := exec.Command("ping", "-c", "1", vm.Address)
	if err := cmd.Run(); err != nil {
		fmt.Printf("error pinging %s: %s", vm.Address, err)
		return false
	}
	return true
}
