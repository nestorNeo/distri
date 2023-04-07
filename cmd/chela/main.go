package main

import (
	"github.com/nestorneo/distri/apis"
	"github.com/nestorneo/distri/nodos"
)

func main() {
	vecinos := []nodos.Nodo{
		{
			Addr: "localhost:9002",
			Name: "tequila",
		},
		{
			Addr: "localhost:9003",
			Name: "vodka",
		},
	}

	r := apis.GetRouterApp("chela", vecinos)
	r.Run("0.0.0.0:9001")
}
