package core

import (
	"context"
	"errors"
	"fmt"

	"log/slog"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"gorm.io/gorm"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
)

type PostService interface {
	UpdatePost(ctx context.Context, postID uint, input model.Post) (model.Post, error)
	CreatePost(ctx context.Context, req dto.CreatePostRequest) (*model.Post, error)
	GetByID(ctx context.Context, id uint) (*dto.PostResponseDTO, error)
	DeletePost(ctx context.Context, id uint) (*model.Post, error)
}

type postService struct {
	repo repository.PostRepo
}

func NewPostService(repo repository.PostRepo) PostService {
	return &postService{
		repo: repo,
	}
}

func (s *postService) CreatePost(ctx context.Context, req dto.CreatePostRequest) (*model.Post, error) {

	if req.JobID != nil && *req.JobID == 0 {
		return nil, errors.New("invalid job id")
	}

	var postImages []model.Image
    for _, url := range req.Images {
        if url != "" {
            postImages = append(postImages, model.Image{
                URL: url,
            })
        }
    }
	post := model.Post{

		Title:     req.Title,
		UserID:    req.UserID,
		JobID:     req.JobID,
		CompanyID: req.CompanyID,
		Content:   req.Content,
		Images:    postImages,
	}
	if post.JobID != nil {
		job, err := s.repo.GetJobByID(ctx, *post.JobID)
		if err != nil {
			slog.Error("[PostService] Job no encontrado", "jobID", *post.JobID, "error", err)
			return nil, fmt.Errorf("la vacante con ID %d no existe", *post.JobID)
		}

		if job.Status != string(model.On) {
			return nil, fmt.Errorf("no se puede publicar: la vacante está %s", job.Status)
		}
		if post.CompanyID != nil && *post.CompanyID != job.CompanyID {
			return nil, fmt.Errorf("la empresa del post no coincide con la empresa de la vacante")
		}
		post.CompanyID = &job.CompanyID
	}
	err := s.repo.CreatePost(ctx, &post)

	if err != nil {
		slog.Error("[PostService][CreatePost] unable create post", "error", err)
		return nil, err
	}
	return &post, nil
}

func (s *postService) UpdatePost(ctx context.Context, postID uint, p model.Post) (model.Post, error) {

	existingPost, err := s.repo.GetByID(ctx, postID)
	if err != nil {
		return model.Post{}, errors.New("post no encontrado")
	}

	p.UserID = existingPost.UserID
	p.CreatedAt = existingPost.CreatedAt

	return s.repo.EditPost(ctx, postID, p)
}

func (s *postService) GetByID(ctx context.Context, id uint) (*dto.PostResponseDTO, error) {

	post, err := s.repo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("The post not found")
		}
		return nil, err
	}
	response := utils.MapToCleanPost(post)

	return response, nil
}

func (s *postService) DeletePost(ctx context.Context, id uint) (*model.Post, error) {

	if s.repo.IsAlreadyDeleted(ctx, id) {
		slog.Error("The post is already Deleted")
		return nil, errors.New("The post is already deleted")
	}
	post, err := s.repo.DeletePost(ctx, id)

	if err != nil {
		slog.Error("[PostService][CreatePost] Unable to delete post", "error", err)
		return nil, err
	}
	return post, nil
}
