package main

import (
	"fmt"

	"github.com/garethpaul/purpleair-go"
)

func main() {
	client := purpleair.NewClient()
	s := client.Sensor("68335")
	for i := 1; i < len(s.Results); i++ {
		fmt.Println("Air Quality: " + s.Results[i].PM25Value)
	}
}
