package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

type Config struct {
	VMAddresses  map[string]string   `toml:"vm_addresses"`
	Dependencies map[string][]string `toml:"dependencies"`
}

func (c *Config) LoadFromFile(path string) error {
	_, err := toml.DecodeFile(path, c)
	if err != nil {
		log.Printf("Error decoding file: %s", err.Error())
		return err
	}
	log.Printf("Successfully decoded file: %s", path)
	return nil
}

func main() {
	var config Config
	log.Printf("config: %+v\n", config)
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}

	var vmAddresses []string
	for _, address := range config.VMAddresses {
		vmAddresses = append(vmAddresses, address)
	}

	type vmData struct {
		Status       string
		Dependencies []string
	}

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"ternary": func(b bool, t, f interface{}) interface{} {
			if b {
				return t
			}
			return f
		},
	})

	r.GET("/", func(c *gin.Context) {
		vmStatus := make(map[string]vmData)
		log.Println("Looping over VM Addresses...")
		for _, address := range vmAddresses {
			log.Printf("Checking VM status for %s...", address)
			status := CheckVMStatus(address, 5*time.Second)
			dependencies, ok := config.Dependencies[address]
			if !ok {
				dependencies = []string{}
			}
			vmStatus[address] = vmData{Status: status, Dependencies: dependencies}
		}
		log.Println("Finished looping over VM addresses.")

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"vmStatus": vmStatus,
		})
	})

	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*.tmpl")

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func CheckVMStatus(address string, timeout time.Duration) string {
	log.Printf("Checking VM status for %s", address)
	conn, err := net.DialTimeout("tcp", address+":22", timeout)
	if err != nil {
		log.Printf("Error checking VM Status for %s: %s", address, err)
		return "offline"
	}
	defer conn.Close()
	log.Printf("VM %s is online", address)
	return "online"
}

func ternary(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		log.Printf("Condition was true, returning %v", trueValue)
		return trueValue
	}
	log.Printf("Condition was false, returning %v", falseValue)
	return falseValue
}
