package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	arg := len(os.Args)

	if arg <= 1 {
		fmt.Println("error input!")
		return
	}

	for i := 1; i < arg; i++ {
		list := strings.Split(os.Args[i], ".")
		if len(list) != 4 {
			fmt.Println("not IPv4 address!")
			return
		}
		var cur []string
		for j := 0; j < 4; j++ {
			a, err := strconv.ParseInt(list[j], 10, 0)
			if err != nil {
				fmt.Println("not IPv4 address ! ")
				return
			}
			if a > 255 || a < 0 {
				fmt.Println("not IPv4 address ! ")
				return
			}

			x := strconv.FormatInt(a, 16)
			if len(x) < 2 {
				x = "0" + x
			}
			cur = append(cur, strings.ToUpper(x))
		}

		fmt.Printf("IPv4 : %s \n", os.Args[i])
		fmt.Printf("IPv6 : 64:FF9B:0:0:0:0:%s%s:%s%s \n", cur[0], cur[1], cur[2], cur[3])
	}

}
