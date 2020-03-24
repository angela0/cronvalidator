package main

import (
	"flag"
	"fmt"
	"github.com/angela0/cronvalidator"
	"os"
)

var (
	spec string
)

func main() {

	flag.Usage = func() {
		fmt.Printf("usage: %s <cron>\n", os.Args[0])
	}
	flag.Parse()

	switch len(os.Args) {
	case 2:
		spec = os.Args[1]
	default:
		flag.Usage()
		return
	}

	if err := cronvalidator.Validate(spec); err != nil {
		fmt.Printf("%s has error: %v\n", spec, err)
		return
	}
	fmt.Printf("%s is ok\n", spec)
}
