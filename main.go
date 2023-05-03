package main

import (
	"fmt"
	"net/http"
	"sync"

	CheckVMStatus "systemChecker/pkg/CheckVMStatus"
	vmStructs "systemChecker/pkg/structs"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var vms []vmStructs.VM
	if err := viper.UnmarshalKey("vms", &vms); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	var mu sync.Mutex

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

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

	// Define routes
	r.GET("/", func(c *gin.Context) {
		var wg sync.WaitGroup

		for i := range vms {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				vms[i].Status = CheckVMStatus.CheckVMStatus(vms[i])
			}(i)
		}

		wg.Wait()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"vms": vms,
		})
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
