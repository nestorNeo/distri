package apis

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestorneo/distri/middleware"
	"github.com/nestorneo/distri/nodos"
	"github.com/nestorneo/distri/respuestas"
)

func GetRouterApp(msg string, vecinos []nodos.Nodo) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GuidMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
		go func() {
			for _, vecino := range vecinos {
				resp, _ := http.Get("http://" + vecino.Addr + "/ping")
				if resp != nil {
					LeerRespuesta(resp)
				}
			}
		}()
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
