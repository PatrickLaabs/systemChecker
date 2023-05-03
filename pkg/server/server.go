package server

import (
	"net/http"
	"sync"
	CheckVMStatus "systemChecker/pkg/CheckVMStatus"
	vmStructs "systemChecker/pkg/structs"

	"github.com/gin-gonic/gin"
)

func Server() {
	var vms []vmStructs.VM

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

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

	r.Run()
}
