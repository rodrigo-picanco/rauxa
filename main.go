package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
         "gopkg.in/yaml.v3"
        "os"
)

type Place struct {
        Name        string
        Neighbourhood string
        Cuisine     string
}

func main() {
        var places []Place
         data, err := os.ReadFile("places.yml")
        if err != nil {
                panic(err)
        }

        if err = yaml.Unmarshal([]byte(data), &places); err != nil {
                panic(err)
        }



        r := gin.Default()
        r.LoadHTMLGlob("templates/**")
        r.GET("/places", func(c *gin.Context) {
                c.HTML(200, "places.tmpl", gin.H{
                        "places": places,
                })
        })

        r.GET("/places/:id", func(c *gin.Context) {
                id := c.Param("id")
                i, e := strconv.Atoi(id)

                if e != nil {
                        c.String(404, "Not found")
                        return
                }

                place := places[i]

                c.HTML(200, "place.tmpl", gin.H{
                        "Name": place.Name,
                        "Neighbourhood": place.Neighbourhood,
                        "Cuisine": place.Cuisine,
                })
        })

        r.Static("/assets", "./assets")

        r.Run(":8088")
}
