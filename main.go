package main

import (
	"fmt"
	"time"
)

var slaveDns = map[int]map[string]interface{}{
	0: {"connectstring": "host0", "weight": 2},
	1: {"connectstring": "host1", "weight": 4},
	2: {"connectstring": "host2", "weight": 7},
	3: {"connectstring": "host3", "weight": 9},
}

var i int = -1  //表示上一次选择的服务器
var cw int = 0  //表示当前调度的权值
var gcd int = 1 //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2

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

//获取最大权重值
func getMaxWeight() int {
	max := 0
	for _, v := range slaveDns {
		if weight, _ := v["weight"].(int); weight >= max {
			max = weight
		}
	}
	return max
}

func main() {
	note := map[string]int{}
	s_time := time.Now().Unix()

	for i := 0; i < 40; i++ {
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
