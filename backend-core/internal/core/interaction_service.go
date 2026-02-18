package core

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

type InteractionRepo interface {
	GetReactionsByPost(postID uint) ([]model.Interaction, error)
	TogglePostReaction(userID string, postID uint, reactionID int) (*model.Interaction, string, error)
}

type interactionService struct {
	repo repository.InteractionRepo
}

func NewInteractionService(repo repository.InteractionRepo) InteractionRepo {
	return &interactionService{repo: repo}
}

func (s *interactionService) TogglePostReaction(userID string, postID uint, reactionID int) (*model.Interaction, string, error) {
	return s.repo.TogglePostReaction(userID, postID, reactionID)
}

func (s *interactionService) GetReactionsByPost(postID uint) ([]model.Interaction, error) {
	return s.repo.GetReactionsByPost(postID)
}
