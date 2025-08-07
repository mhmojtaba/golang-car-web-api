package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	r.GET("/", handler.TestHandler)
	r.POST("/userById/:id", handler.UserById)

	r.GET("/head1", handler.HeaderBinderMethod1)
	r.GET("/head2", handler.HeaderBinderMethod2)
	r.GET("/head3", handler.HeaderBinderMethod3)
	r.GET("/head4", handler.HeaderBinderMethod4)

	r.GET("/query", handler.ReadQuery)

	r.GET("/uri/:code/:id", handler.UriBinder)

	r.POST("/body", handler.BodyBinder)

	r.POST("/form", handler.FormBinder)

	r.POST("/file", handler.FileBinder)
}
