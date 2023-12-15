package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikaelhg/comments-api/internal/services"
)

//go:generate mockgen -package=controllers -destination=mocks_generated.go . CommentController
type CommentController interface {
	GetComment(c echo.Context) error
}

type CommentControllerImpl struct {
	CommentController
	commentService services.CommentService `di.inject:"commentService"`
}

func (cc *CommentControllerImpl) GetComment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	comment, err := cc.commentService.GetComment(id)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, comment)
	return nil
}
