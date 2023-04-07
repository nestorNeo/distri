package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nestorneo/distri/respuestas"
)

func main() {

	for i := 0; i < 100; i++ {
		resp, err := http.Get("http://localhost:9001/ping")
		if err != nil {
			log.Println(err)
			log.Println("stopping")
			break
		}

		switch resp.StatusCode {
		case 200:
			log.Println("old good !!!")
			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Panicln("ERROR no se puede leer cuerpo de respuesta", err)
				continue
			}

			respuesta := respuestas.Respuesta{}
			if err = json.Unmarshal(bytes, &respuesta); err != nil {
				log.Println("no se pudo pasar a variable la respuesta")
			}

			log.Println(respuesta)

		default:
			log.Println("no sabemos que paso ", resp.StatusCode)
		}
	}

}
