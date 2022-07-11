package main

import (
	"monitor/internal/database"
	"monitor/models"
	"github.com/gin-gonic/gin"
	// "encoding/json"
	// "fmt"
)

func main() {

	// jsonStr := `{"iron_info":{"account_pubkey":"","miner_count":"0","node_count":"0","node_graffi":"","node_height":"","node_version":""},"system_info":{"cpu_cores":"8","cpu_usage":"1.96%","ip_address":"172.16.200.12","mem_free":"56.72G","mem_total":"67.33G","mem_usage":"1.87%","mem_used":"1.26G","pc_hostname":"xw127","pc_os":"linux","pc_platform":"ubuntu-20.04"},"update_time":"2022-06-21 16:07:55"}`
	// var show Collections
    // err := json.Unmarshal([]byte(jsonStr), &show)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }

	router := gin.Default()

	// router.GET("/show", func(c *gin.Context) {
	// 	c.IndentedJSON(200, show)
	// })

	var param models.Collections

	router.POST("/postjson", func(c *gin.Context) {
		// param := make(map[string]interface{})
		param = models.Collections{}
		err := c.BindJSON(&param)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
      		return
		}
		// fmt.Println(param)
		// fmt.Printf("%T\n", param)
		// fmt.Printf("%T\n", string(param))
		database.PostIronMysql(param)
		c.JSON(200, gin.H{
			"results": param,
		})
	})

	
	// curl -X POST http://127.0.0.1:9100/postjson -d '{"account_pubkey":"","cpu_cores":"8","cpu_usage":"1.00%","ip_address":"172.16.200.12","mem_total":"67.33G","mem_usage":"1.90%","miner_count":"0","node_count":"0","node_graffi":"","node_height":"","node_version":"","update_time":"2022-06-22 16:00:58"}'
	
	router.Run(":9100")
}