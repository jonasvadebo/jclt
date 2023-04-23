package main

import (
	"flag"
	"fmt"
)

var command = flag.String("command", "deploy", "Default deployment.")

var strategy bool

func init() {
	flag.BoolVar(&strategy, "recreate", false, "Use recreate strategy")
	flag.BoolVar(&strategy, "r", false, "Use recreate strategy")
}

func main() {
	flag.Parse()

	if strategy == true {
		fmt.Printf("%s with deploy strategy: recreate\n", *command)
	} else {
		fmt.Printf("%s with deploy strategy: rolling\n", *command)
	}
}
