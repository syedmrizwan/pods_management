package api

import (
	"fmt"
	"github.com/syedmrizwan/pods_management/model"
	"log"
	"time"
)

var configs = []model.ConfigurationInfo{
	{IpAddress: "10.2.0.1", TypeName: "sdwan", DatastoreID: 123, ClusterID: 44, Password: "pass", UserName: "uuss", VcenterId: 123},
	{IpAddress: "10.2.0.2", TypeName: "sdwan", DatastoreID: 123, ClusterID: 44, Password: "pass", UserName: "uuss", VcenterId: 123},
	{IpAddress: "10.2.0.3", TypeName: "sdwan", DatastoreID: 123, ClusterID: 44, Password: "pass", UserName: "uuss", VcenterId: 123},
	{IpAddress: "10.2.0.4", TypeName: "sdwan", DatastoreID: 123, ClusterID: 44, Password: "pass", UserName: "uuss", VcenterId: 123},
}

var slaveDns = map[int]map[string]interface{}{
	0: {"connectstring": configs[0], "weight": 1},
	1: {"connectstring": configs[1], "weight": 1},
	2: {"connectstring": configs[2], "weight": 1},
}

var i int = -1  // indicates the last selected server
var cw int = 0  // indicates the weight of the current schedule
var gcd int = 1 //The greatest common divisor of the current weight of ownership. For example, the greatest common divisor of 2,4,8 is: 2
/*
 Algorithm ideas:
 Cw is the current weight, traversing each server. If the weight of the server is greater than cw, the server executes once, after a round of polling, cw-gcd;
 Repeat the appeal step
*/
func getDns() string {
	for {
		i = (i + 1) % len(slaveDns)
		if i == 0 {
			cw = cw - gcd
			if cw <= 0 {
				cw = getMaxWeight()
				if cw == 0 {
					return ""
				}
			}
		}

		if weight, _ := slaveDns[i]["weight"].(int); weight >= cw {
			res := slaveDns[i]["connectstring"]
			config, ok := res.(model.ConfigurationInfo)
			if !ok {
				log.Printf("got data of type %T but wanted int", res)
				return "Error"
			}
			return config.IpAddress
		}
	}
}

func getMaxWeight() int {
	max := 0
	for _, v := range slaveDns {
		if weight, _ := v["weight"].(int); weight >= max {
			max = weight
		}
	}

	return max
}

func TestRoundRobin() {

	note := map[string]int{}

	s_time := time.Now().Unix()

	for i := 0; i < 100; i++ {
		s := getDns()
		fmt.Println(s)
		if note[s] != 0 {
			note[s]++
		} else {
			note[s] = 1
		}
	}

	e_time := time.Now().Unix()

	fmt.Println("total time: ", e_time-s_time)

	fmt.Println("--------------------------------------------------")

	for k, v := range note {
		fmt.Println(k, " ", v)
	}
}
