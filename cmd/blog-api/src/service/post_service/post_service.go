package post_service

import "github.com/sum-project/ublog/cmd/blog-api/src/repository/post_repository"

type PostService interface {
	GetPost()
	GetPosts()
	AddPost()
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
func (s *postService) AddPost() {

}
func (s *postService) UpdatePost() {

}
func (s *postService) DeletePost() {

}
