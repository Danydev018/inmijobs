package core

import (
	"errors"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(userID uint, req dto.CreateCommentReq) (*model.Comment, error) {
	newComment := &model.Comment{
		Content: req.Content,
		PostID:  req.PostID,
		UserID:  userID,
	}
	return s.repo.Create(newComment)
}

func (s *CommentService) GetCommentsByPost(postID uint) ([]model.Comment, error) {
	return s.repo.GetByPostID(postID)
}

func (s *CommentService) UpdateComment(userID uint, commentID uint, req dto.UpdateCommentReq) (*model.Comment, error) {
	comment, err := s.repo.GetByID(commentID)
	if err != nil {
		return nil, err
	}
	if comment.UserID != userID {
		return nil, errors.New("no autorizado")
	}
	comment.Content = req.Content
	return s.repo.Update(comment)
}

func (s *CommentService) DeleteComment(userID uint, commentID uint) error {
	comment, err := s.repo.GetByID(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		return errors.New("no autorizado")
	}
	return s.repo.Delete(commentID)
}