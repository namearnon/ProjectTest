package route

import (
	controller "ProjectTest/controller/candidate"

	"github.com/labstack/echo/v4"
)

func TestRoute(e *echo.Echo) {
	c := controller.NewControllerTest()
	e.GET("api/beer", c.BeerSelect)
	e.POST("api/beer", c.BeerInsert)
	e.PUT("api/beer/:id", c.BeerUpdate)
	e.DELETE("api/beer/:id", c.BeerDelete)
	// e.POST("api/JWT", c.PostJWT)
}
