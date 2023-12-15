package main

import (
	"reflect"

	"github.com/goioc/di"
	"github.com/knadh/koanf/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mikaelhg/comments-api/internal/controllers"
	"github.com/mikaelhg/comments-api/internal/dao"
	"github.com/mikaelhg/comments-api/internal/services"
)

func init() {
	_, _ = di.RegisterBean("commentsDao", reflect.TypeOf((*dao.CommentDaoImpl)(nil)))
	_, _ = di.RegisterBean("commentsService", reflect.TypeOf((*services.CommentServiceImpl)(nil)))
	_, _ = di.RegisterBean("commentsController", reflect.TypeOf((*controllers.CommentControllerImpl)(nil)))
	_ = di.InitializeContainer()
}

var commentsController controllers.CommentController

var k = koanf.New(".")

func main() {
	commentsController = di.GetInstance("commentsController").(controllers.CommentController)
	e := echo.New()
	e.Use(middleware.RequestID())
	e.GET("/comments/:id", commentsController.GetComment)
	e.Logger.Fatal(e.Start(":1323"))
}
