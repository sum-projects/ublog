package post_service

import (
	"github.com/sum-project/ublog/cmd/blog-api/src/domain/post"
	"github.com/sum-project/ublog/cmd/blog-api/src/repository/post_repository"
	"github.com/sum-project/ublog/pkg/uerror"
)

type PostService interface {
	GetPost()
	GetPosts()
	AddPost(post.Post) (*post.Post, uerror.Error)
	UpdatePost()
	DeletePost()
}

type postService struct {
	postRepository post_repository.PostRespoitory
}

func NewPostService(postRepository post_repository.PostRespoitory) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (s *postService) GetPost() {

}
func (s *postService) GetPosts() {

}
func (s *postService) AddPost(p post.Post) (*post.Post, uerror.Error) {
	return s.postRepository.Create(p)
}
func (s *postService) UpdatePost() {

}
func (s *postService) DeletePost() {

}
