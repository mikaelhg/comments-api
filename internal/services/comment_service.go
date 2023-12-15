package services

import "github.com/mikaelhg/comments-api/internal/domain"

//go:generate mockgen -package=services -destination=mocks_generated.go . CommentService
type CommentService interface {
	GetComment(id int64) (*domain.Comment, error)
}

type CommentServiceImpl struct {
	CommentService
}

func (cs *CommentServiceImpl) GetComment(id int64) (*domain.Comment, error) {
	return &domain.Comment{
		CommentId: id,
		ThreadId:  1234,
		Sender:    "foo",
		Text:      "bar",
	}, nil
}
