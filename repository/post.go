package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	Update(id int, post *model.Post) error
	Delete(id int) error
	GetByID(id int) (*model.Post, error)
	GetBySlug(slug string) (*model.Post, error)
	GetAll(limit, offset int) ([]model.Post, error)
	GetPublished(limit, offset int) ([]model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Update(id int, post *model.Post) error {
	if err := r.db.Model(&model.Post{}).
		Where("id = ?", id).
		Updates(post).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *postRepository) Delete(id int) error {
	return r.db.Delete(&model.Post{}, id).Error
}

func (r *postRepository) GetByID(id int) (*model.Post, error) {
	var post model.Post
	err := r.db.Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetBySlug(slug string) (*model.Post, error) {
	var post model.Post
	err := r.db.Where("slug = ?", slug).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll(limit, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *postRepository) GetPublished(limit, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.
		Where("published = ?", true).
		Order("published_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error
	return posts, err
}