package main

import (
	"github.com/enkhalifapro/order-packs-calculator/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
