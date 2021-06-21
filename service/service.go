package service

import (
	"bzz-collector/api"
	"fmt"
)

func StartService() {
	//r := gin.Default()
	//r.GET("/bzz-info", func(c *gin.Context) {
	ContainersInfoResults := api.GetContainersInspec()
	//	c.JSON(200, gin.H{
	//		"data": ContainersInfoResults,
	//	})
	//})
	//apiServer := fmt.Sprintf(":%s", port)
	//r.Run(apiServer)
	for _, val := range ContainersInfoResults {
		fmt.Printf("%v\n", val)
	}

}
