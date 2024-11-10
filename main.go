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
                c.HTML(200, "places.tpl", gin.H{
                        "places": sortPlaces(places, c.DefaultQuery("sort", "name")),
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

                c.HTML(200, "place.tpl", gin.H{
                        "Name": place.Name,
                        "Neighbourhood": place.Neighbourhood,
                        "Cuisine": place.Cuisine,
                })
        })

        r.Static("/assets", "./assets")

        r.Run(":8088")
}

func sortPlaces(places []Place, sort string) []Place {
        for i := 0; i < len(places); i++ {
                for j := i + 1; j < len(places); j++ {
                        if sort == "name" && places[i].Name > places[j].Name {
                                places[i], places[j] = places[j], places[i]
                        } else if sort == "neighbourhood" && places[i].Neighbourhood > places[j].Neighbourhood {
                                places[i], places[j] = places[j], places[i]
                        } else if sort == "cuisine" && places[i].Cuisine > places[j].Cuisine {
                                places[i], places[j] = places[j], places[i]
                        }
                }
        }
        return places
}

