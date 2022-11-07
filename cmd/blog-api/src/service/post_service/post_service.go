package post_service

import (
	"github.com/google/uuid"
	"github.com/sum-project/ublog/cmd/blog-api/src/domain/post"
	"github.com/sum-project/ublog/cmd/blog-api/src/repository/post_repository"
	"github.com/sum-project/ublog/pkg/uerror"
)

type PostService interface {
	GetPost(uuid.UUID) (*post.Post, uerror.Error)
	GetPosts() (post.Posts, uerror.Error)
	AddPost(post.Post) (*post.Post, uerror.Error)
	UpdatePost(uuid.UUID, post.Post) (*post.Post, uerror.Error)
	DeletePost(uuid.UUID) uerror.Error
}

type postService struct {
	postRepository post_repository.PostRespoitory
}

func NewPostService(postRepository post_repository.PostRespoitory) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (s *postService) GetPost(id uuid.UUID) (*post.Post, uerror.Error) {
	return s.postRepository.Get(id)
}
func (s *postService) GetPosts() (post.Posts, uerror.Error) {
	return s.postRepository.GetAll()
}
func (s *postService) AddPost(p post.Post) (*post.Post, uerror.Error) {
	return s.postRepository.Create(p)
}
func (s *postService) UpdatePost(id uuid.UUID, p post.Post) (*post.Post, uerror.Error) {
	return s.postRepository.Update(id, p)
}
func (s *postService) DeletePost(id uuid.UUID) uerror.Error {
	return s.postRepository.Delete(id)
}
