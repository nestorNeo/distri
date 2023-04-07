package apis

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestorneo/distri/middleware"
	"github.com/nestorneo/distri/nodos"
	"github.com/nestorneo/distri/respuestas"
	"github.com/nestorneo/distri/solicitudes"
)

func GetRouterApp(msg string, vecinos map[string]nodos.Nodo) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GuidMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
		go func() {

			if vecinos == nil {
				return
			}

			if vodka, ok := vecinos["vodka"]; ok {
				resp, _ := http.Get("http://" + vodka.Addr + "/ping")
				if resp != nil {
					LeerRespuesta(resp)
				}
			}
			if vodka, ok := vecinos["tequila"]; ok {

				ordenTequila := solicitudes.BuzonTequila{
					Margarita:     true,
					EnLasRocas:    false,
					Instrucciones: "para chela xD please",
				}

				jsonData, err := json.Marshal(ordenTequila)

				if err != nil {
					log.Println(err)
					log.Println("error mandandole info a tequila")
					return
				}

				resp, err := http.Post("http://"+vodka.Addr+"/buzon",
					"application/json",
					bytes.NewBuffer(jsonData),
				)

				if err != nil {
					log.Println("HEY TEQUILA", err)
				}

				if resp != nil {
					LeerRespuesta(resp)
				}
			}
		}()
	})

	return r
}

func GetRouterTequila(msg string, vecinos map[string]nodos.Nodo) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GuidMiddleware())

	r.POST("/buzon", func(c *gin.Context) {
		log.Println("con el tequila")

		var ordenDeTequila solicitudes.BuzonTequila

		if err := c.BindJSON(&ordenDeTequila); err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": msg,
			"info":    "recibido",
		})

		log.Println(ordenDeTequila)

	})

	return r
}

func LeerRespuesta(resp *http.Response) {
	switch resp.StatusCode {
	case 200:
		log.Println("old good !!!")
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panicln("ERROR no se puede leer cuerpo de respuesta", err)
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
