package advertisement

import (
	"ProjectTest/middlewares/core"
	"ProjectTest/model"
	service "ProjectTest/service/candidate"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

type ControllerTest struct {
}

func NewControllerTest() *ControllerTest {
	return &ControllerTest{}
}

func (*ControllerTest) BeerSelect(c echo.Context) error {
	cc := c.(core.IContext)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
	}
	postData := model.PostData{}
	json.Unmarshal(body, &postData)
	data := service.BeerSer(cc).GetBeer(&postData)
	return c.JSON(200, data)
}

func (*ControllerTest) BeerInsert(c echo.Context) error {
	cc := c.(core.IContext)
	file, err := c.FormFile("myFile")
	if err != nil {
		log.Println(err)
	}
	beerName := c.FormValue("beerName")
	beerType := c.FormValue("beerType")
	BeerDesc := c.FormValue("beerDesc")
	postData := model.GetBeerData{BeerName: beerName, BeerType: beerType, BeerDesc: BeerDesc, BeerImage: file}

	// json.Unmarshal(body, &postData)
	data := service.BeerSer(cc).PostBeer(&postData)
	return c.JSON(201, data)
}

func (*ControllerTest) BeerUpdate(c echo.Context) error {
	cc := c.(core.IContext)
	file, err := c.FormFile("myFile")
	if err != nil {
		log.Println(err)
	}
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(405, map[string]string{
			"message": "โปรดแกรมกรอกไอดี",
		})
	}
	beerName := c.FormValue("beerName")
	beerType := c.FormValue("beerType")
	BeerDesc := c.FormValue("beerDesc")
	postData := model.GetBeerData{ID: id, BeerName: beerName, BeerType: beerType, BeerDesc: BeerDesc, BeerImage: file}

	// json.Unmarshal(body, &postData)
	data := service.BeerSer(cc).PutBeer(&postData)
	return c.JSON(201, data)
}

func (*ControllerTest) BeerDelete(c echo.Context) error {
	cc := c.(core.IContext)
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(405, map[string]string{
			"message": "โปรดแกรมกรอกไอดี",
		})
	}
	postData := model.GetBeerData{ID: id}
	data := service.BeerSer(cc).DeleteBeer(&postData)
	return c.JSON(200, data)
}
