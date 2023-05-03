package configloader

import (
	"fmt"

	"github.com/spf13/viper"

	vmStructs "systemChecker/pkg/structs"
)

func Configloader() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var vms []vmStructs.VM
	if err := viper.UnmarshalKey("vms", &vms); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
