package api

import (
	"fmt"
	"time"
)

var slaveDns = map[int]map[string]interface{}{
	0: {"connectstring": "root@tcp(172.16.0.164:3306)/shiqu_tools?charset=utf8", "weight": 8},
	1: {"connectstring": "root@tcp(172.16.0.165:3306)/shiqu_tools?charset=utf8", "weight": 4},
	2: {"connectstring": "root@tcp(172.16.0.166:3306)/shiqu_tools?charset=utf8", "weight": 2},
}

var i int = -1  // indicates the last selected server
var cw int = 0  // indicates the weight of the current schedule
var gcd int = 2 //The greatest common divisor of the current weight of ownership. For example, the greatest common divisor of 2,4,8 is: 2
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
			return slaveDns[i]["connectstring"].(string)
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
