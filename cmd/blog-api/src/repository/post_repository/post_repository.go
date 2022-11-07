package post_repository

import (
	"github.com/google/uuid"
	"github.com/sum-project/ublog/cmd/blog-api/src/clients/mysql"
	"github.com/sum-project/ublog/cmd/blog-api/src/domain/post"
	"github.com/sum-project/ublog/pkg/uerror"
)

const (
	queryGetPost    = "SELECT id, title, content, created_at, updated_at, deleted_at FROM posts WHERE id=?;"
	queryGetAllPost = "SELECT id, title, content, created_at, updated_at, deleted_at FROM posts;"
	queryInsertPost = "INSERT INTO posts(id, title, content) VALUES(?, ?, ?);"
	queryUpdatePost = "UPDATE posts SET title=?, content=? WHERE id=?;"
	queryDeletePost = "UPDATE posts SET deleted_at=NOW() WHERE id=?;"
)

type PostRespoitory interface {
	Get(uuid.UUID) (*post.Post, uerror.Error)
	GetAll() (post.Posts, uerror.Error)
	Create(post.Post) (*post.Post, uerror.Error)
	Update(uuid.UUID, post.Post) (*post.Post, uerror.Error)
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

	p.ID = uuid.New()
	if _, err = stmt.Exec(p.ID.String(), p.Title, p.Content); err != nil {
		return nil, uerror.NewInternalServerError("error when trying to insert post", err)
	}

	return repo.Get(p.ID)
}

func (repo *postRepository) Update(id uuid.UUID, p post.Post) (*post.Post, uerror.Error) {
	stmt, err := mysql.Client.Prepare(queryUpdatePost)
	if err != nil {
		return nil, uerror.NewInternalServerError("error when trying to prepare update post statement", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Title, p.Content, id.String())
	if err != nil {
		return nil, uerror.NewInternalServerError("error when trying to update post", err)
	}

	return repo.Get(id)
}

func (repo *postRepository) Delete(id uuid.UUID) uerror.Error {
	stmt, err := mysql.Client.Prepare(queryDeletePost)
	if err != nil {
		return uerror.NewInternalServerError("error when trying to prepare delete post statement", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id.String()); err != nil {
		return uerror.NewInternalServerError("error when tying to delete post", err)
	}

	return nil
}
