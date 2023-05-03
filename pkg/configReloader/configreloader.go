package configreloader

import (
	"fmt"
	"sync"

	vmStructs "systemChecker/pkg/structs"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Configreloader() {
	var mu sync.Mutex
	var vms []vmStructs.VM

	fmt.Printf("Initial config: %#v\n", vms)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
			return
		}

		var newVms []vmStructs.VM
		if err := viper.UnmarshalKey("vms", &newVms); err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
			return
		}
		fmt.Printf("Reloaded config: %#v\n", newVms)

		mu.Lock()
		vms = newVms
		mu.Unlock()
	})
}
