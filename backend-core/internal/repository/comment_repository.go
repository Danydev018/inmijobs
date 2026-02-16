package repository

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

// CommentRepository DEBE tener estos métodos exactos para que el Service no falle
type CommentRepository interface {
	Create(comment *model.Comment) (*model.Comment, error)
	GetByPostID(postID uint) ([]model.Comment, error)
	GetByID(id uint) (*model.Comment, error)               // Esto quita el error de "GetByID undefined"
	Update(comment *model.Comment) (*model.Comment, error) // Esto arregla el "WrongArgCount"
	Delete(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *model.Comment) (*model.Comment, error) {
	err := r.db.Create(comment).Error
	return comment, err
}

func (r *commentRepository) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) GetByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) Update(comment *model.Comment) (*model.Comment, error) {
	err := r.db.Save(comment).Error
	return comment, err
}

func (r *commentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Comment{}, id).Error
}
