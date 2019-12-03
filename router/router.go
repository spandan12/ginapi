package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spandan12/ginapi/api"
)



func New() *gin.Engine{
	r := gin.Default()
	// fmt.Println(reflect.TypeOf(r))
	rg := r.Group("/api")
	{
		rg.POST("/", api.Create)
		rg.GET("/", api.FetchAll)
		rg.GET("/:id", api.Fetch)
		rg.PUT("/:id", api.Update)
		rg.DELETE("/:id", api.Delete)
	}

	return r
}