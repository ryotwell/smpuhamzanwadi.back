package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	Update(slug string, post *model.Post) error
	Delete(slug string) error
	GetByID(id int) (*model.Post, error)
	GetBySlug(slug string) (*model.Post, error)
	GetAll(limit, page int, q string) ([]model.Post, error)
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

func (r *postRepository) Update(slug string, post *model.Post) error {
	return r.db.Model(&model.Post{}).
		Where("slug = ?", slug). // gunakan slug
		Updates(post).
		Error
}

func (r *postRepository) Delete(slug string) error {
	return r.db.
		Where("slug = ?", slug). // nanti akan diganti pada service
		Delete(&model.Post{}).
		Error
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

func (r *postRepository) GetAll(limit, page int, q string) ([]model.Post, error) {
	var posts []model.Post
	db := r.db
	offset := (page - 1) * limit

	if q != "" {
		db = db.Where("title ILIKE ? OR content ILIKE ?", "%"+q+"%", "%"+q+"%")
	}

	err := db.
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
