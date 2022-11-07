package post_repository

import (
	"github.com/google/uuid"
	"github.com/sum-project/ublog/cmd/blog-api/src/clients/mysql"
	"github.com/sum-project/ublog/cmd/blog-api/src/domain/post"
	"github.com/sum-project/ublog/pkg/uerror"
)

const (
	queryGetPost    = "SELECT id, title, content, created_at, updated_at, deleted_at FROM users WHERE id=?;"
	queryGetAllPost = "SELECT id, title, content, created_at, updated_at, deleted_at FROM users;"
	queryInsertPost = "INSERT INTO users(title, content) VALUES(?, ?);"
	queryUpdatePost = "UPDATE users SET title=?, content=? WHERE id=?;"
	queryDeletePost = "UPDATE users SET deleted_at=NOW() WHERE id=?;"
)

type PostRespoitory interface {
	Get(uuid.UUID) (*post.Post, uerror.Error)
	GetAll() (post.Posts, uerror.Error)
	Create(post.Post) (*post.Post, uerror.Error)
	Update(post.Post, uuid.UUID) uerror.Error
	Delete(uuid.UUID) uerror.Error
}

type postRepository struct{}

func NewPostRepository() PostRespoitory {
	return &postRepository{}
}

func (repo *postRepository) Get(id uuid.UUID) (*post.Post, uerror.Error) {
	stmt, err := mysql.Client.Prepare(queryGetPost)
	if err != nil {
		return nil, uerror.NewInternalServerError("error when trying to prepare get post statement", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var result post.Post
	if err = row.Scan(&result.ID, &result.Title, &result.Content, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt); err != nil {
		return nil, uerror.NewInternalServerError("error when trying to get post", err)
	}

	return &result, nil
}

func (repo *postRepository) GetAll() (post.Posts, uerror.Error) {
	stmt, err := mysql.Client.Prepare(queryGetAllPost)
	if err != nil {
		return nil, uerror.NewInternalServerError("error when trying to prepare get posts statement", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, uerror.NewInternalServerError("error when trying to get posts", err)
	}

	result := make(post.Posts, 0)

	for rows.Next() {
		var p post.Post
		if err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			return nil, uerror.NewInternalServerError("error when trying to get post", err)
		}
		result = append(result, p)
	}

	return result, nil
}

func (repo *postRepository) Create(p post.Post) (*post.Post, uerror.Error) {
	stmt, err := mysql.Client.Prepare(queryInsertPost)
	if err != nil {
		return nil, uerror.NewInternalServerError("error when trying to prepare insert post statement", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(p.Title, p.Content)
	if err = row.Scan(&p.ID); err != nil {
		return nil, uerror.NewInternalServerError("error when trying to insert post", err)
	}

	var id uuid.UUID
	if err = row.Scan(&id); err != nil {
		return nil, uerror.NewInternalServerError("error when trying to get post uuid", err)
	}

	return &p, nil
}

func (repo *postRepository) Update(p post.Post, id uuid.UUID) uerror.Error {
	return nil
}

func (repo *postRepository) Delete(id uuid.UUID) uerror.Error {
	return nil
}
