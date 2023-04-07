package main

import (
	"github.com/nestorneo/distri/apis"
	"github.com/nestorneo/distri/nodos"
)

func main() {
	// 192.168.100.95 - tequila
	// 192.168.100.33 - vodka
	vecinos := map[string]nodos.Nodo{
		"tequila": {
			Addr: "192.168.100.95:9002",
			Name: "tequila",
		},
		"vodka": {
			Addr: "192.168.100.33:9003",
			Name: "vodka",
		},
	}

	r := apis.GetRouterApp("chela", vecinos)
	r.Run("0.0.0.0:9001")
}
