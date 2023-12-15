package main

import (
	"reflect"

	"github.com/goioc/di"
	"github.com/labstack/echo/v4"
	"github.com/mikaelhg/comments-api/internal/controllers"
	"github.com/mikaelhg/comments-api/internal/services"
)

func init() {
	_, _ = di.RegisterBean("commentService", reflect.TypeOf((*services.CommentServiceImpl)(nil)))
	_, _ = di.RegisterBean("commentController", reflect.TypeOf((*controllers.CommentControllerImpl)(nil)))
	_ = di.InitializeContainer()
}

func main() {
	e := echo.New()
	e.GET("/comment/:id", di.GetInstance("commentController").(controllers.CommentController).GetComment)
	e.Logger.Fatal(e.Start(":1323"))
}
