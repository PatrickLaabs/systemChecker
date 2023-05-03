package CheckVMStatus

import "os/exec"

type VM struct {
	Name         string
	Address      string
	Status       bool
	Dependencies []string
}

func CheckVMStatus(vm VM) bool {
	cmd := exec.Command("ping", "-c", "1", vm.Address)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
