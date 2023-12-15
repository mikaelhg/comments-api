package dao

//go:generate mockgen -package=dao -destination=mocks_generated.go . CommentDao
type CommentDao interface {
}

type CommentDaoImpl struct {
	CommentDao
}
