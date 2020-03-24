package main

import (
	"flag"
	"fmt"
	"github.com/angela0/cronvalidator"
	"os"
	"strconv"
)

var (
	spec  string = "0 */12 * * *"
	times int    = 8
)

func main() {
	flag.Usage = func() {
		fmt.Printf("usage: %s <cron> <times>\n", os.Args[0])
	}
	flag.Parse()

	switch len(os.Args) {
	case 1:
	case 2:
		spec = os.Args[1]
	case 3:
		spec = os.Args[1]
		times, _ = strconv.Atoi(os.Args[2])
	default:
		flag.Usage()
		return
	}

	ts, err := cronvalidator.Next(spec, times)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("crontab expression: %s, times: %d\n", spec, times)
	for i, t := range ts {
		fmt.Printf("next %d time[s]: %s\n", i+1, t)
	}
}
