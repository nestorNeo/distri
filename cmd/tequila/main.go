package main

import (
	"github.com/nestorneo/distri/apis"
)

func main() {
	r := apis.GetRouterTequila("tequila", nil)
	r.Run("0.0.0.0:9002")
}
