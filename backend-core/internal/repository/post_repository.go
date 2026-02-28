package repository

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepo interface {
	EditPost(ctx context.Context, postID string, p model.Post) (model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) error
	GetByID(ctx context.Context, id string) (*model.Post, error)
	DeletePost(ctx context.Context, id string) (*model.Post, error)
	IsAlreadyDeleted(ctx context.Context, id string) bool
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepo {
	return &postRepository{db: db}
}

func (r *postRepository) GetByID(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post

	err := r.db.WithContext(ctx).
        Preload("User").
        Preload("Company").
		Preload("Company").
        Preload("Job").
        Preload("Images").
		Preload("Interactions").
        Preload("Interactions.User").
        Preload("Interactions.Reaction").
        Preload("Comments.User"). 
		Preload("Comments").
        First(&post,"id = ?", id).Error

	return &post, err
}

func (r *postRepository) EditPost(ctx context.Context, postID string, p model.Post) (model.Post, error) {
	var editedPost model.Post

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&editedPost, "posts.id = ?",postID).Error; err != nil {
			return err
		}

		if err := tx.Model(&editedPost).Updates(p).Error; err != nil {
			return err
		}

		if len(p.Images) > 0 {

			if err := tx.Model(&editedPost).Association("Images").Replace(p.Images); err != nil {
				return err
			}
		}

		if len(p.Comments) > 0 {
			if err := tx.Model(&editedPost).Association("Comments").Replace(p.Comments); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return model.Post{}, err
	}

	res, err := r.GetByID(ctx, postID)
	if err != nil {
		return model.Post{}, err
	}
	return *res, nil
}
func (r *postRepository) CreatePost(ctx context.Context, post *model.Post) error {
	if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (r *postRepository) DeletePost(ctx context.Context, id string) (*model.Post, error) {
	post := model.Post{ID: id}

	if err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Delete(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) IsAlreadyDeleted(ctx context.Context, id string) bool {
	post := model.Post{ID: id}

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return true
	}
	return false
}
