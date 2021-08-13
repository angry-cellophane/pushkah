package main

import (
	"fmt"

	"github.com/angry-cellophane/pushkah/pusher"
)

func main() {
	p := pusher.NewPusher()
	fmt.Print(p)
}
