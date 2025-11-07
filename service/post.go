package service

import (
	"project_sdu/model"
	"project_sdu/repository"
)

type PostService interface {
	CreatePost(post *model.Post) error
	GetAllPosts(limit int, page int, q string) ([]model.Post, error)
	GetPublishedPosts(limit int, offset int) ([]model.Post, error)
	GetPostByID(id int) (*model.Post, error)
	GetPostBySlug(slug string) (*model.Post, error)
	UpdatePost(id int, post *model.Post) error
	DeletePost(id int) error
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo}
}

func (s *postService) CreatePost(post *model.Post) error {
	if err := s.postRepo.Create(post); err != nil {
		return err
	}
	return nil
}

func (s *postService) GetAllPosts(limit int, page int, q string) ([]model.Post, error) {
	posts, err := s.postRepo.GetAll(limit, page, q)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *postService) GetPublishedPosts(limit int, offset int) ([]model.Post, error) {
	posts, err := s.postRepo.GetPublished(limit, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *postService) GetPostByID(id int) (*model.Post, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) GetPostBySlug(slug string) (*model.Post, error) {
	post, err := s.postRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) UpdatePost(id int, post *model.Post) error {
	_, err := s.postRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.postRepo.Update(id, post)
}

func (s *postService) DeletePost(id int) error {
	_, err := s.postRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.postRepo.Delete(id)
}
