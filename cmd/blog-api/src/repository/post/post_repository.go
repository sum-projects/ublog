package post

import (
	"github.com/google/uuid"
	"github.com/sum-project/ublog/cmd/blog-api/src/domain/post"
	"github.com/sum-project/ublog/pkg/uerror"
)

const (
	queryGetPost    = "SELECT id, title, content, created_at, updated_at, deleted_at FROM users WHERE id=?;"
	queryGetAllPost = "SELECT id, title, content, created_at, updated_at, deleted_at FROM users;"
	queryUpdatePost = "UPDATE users SET title=?, content=? WHERE id=?;"
	queryDeletePost = "UPDATE users SET deleted_at=NOW() WHERE id=?;"
)

type PostRespoitory interface {
	Get(uuid.UUID) (post.Post, uerror.Error)
	GetAll() (post.Posts, uerror.Error)
	Create(post.Post) uerror.Error
	Update(post.Post, uuid.UUID) uerror.Error
	Delete(uuid.UUID) uerror.Error
}

type postRepository struct{}

func NewPostRepository() PostRespoitory {
	return &postRepository{}
}

func (repo *postRepository) Get(id uuid.UUID) (post.Post, uerror.Error) {
	return post.Post{}, nil
}

func (repo *postRepository) GetAll() (post.Posts, uerror.Error) {
	return nil, nil
}

func (repo *postRepository) Create(p post.Post) uerror.Error {
	return nil
}

func (repo *postRepository) Update(p post.Post, id uuid.UUID) uerror.Error {
	return nil
}

func (repo *postRepository) Delete(id uuid.UUID) uerror.Error {
	return nil
}
