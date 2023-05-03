package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type VM struct {
	Name    string
	Address string
	Status  bool
}

func main() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var vms []VM
	err = viper.UnmarshalKey("vms", &vms)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// fmt.Println("VMs:")
	// for _, vm := range vms {
	// 	status := checkVMStatus(vm)
	// 	fmt.Printf("%s (%s): %t\n", vm.Name, vm.Address, status)
	// }

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")

	// Define routes
	r.GET("/", func(c *gin.Context) {
		// Check status of each VM
		for i, vm := range vms {
			vms[i].Status = checkVMStatus(vm)
		}

		// Render template with VM information
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"vms": interface{}(vms),
		})
	})

	// Start web server
	r.Run()
}

func checkVMStatus(vm VM) bool {
	cmd := exec.Command("ping", "-c", "1", vm.Address)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
